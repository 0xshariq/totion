package stats

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode"
)

// Statistics holds note statistics
type Statistics struct {
	WordCount      int
	CharCount      int
	LineCount      int
	ParagraphCount int
	ReadingTime    int // in minutes
	SentenceCount  int
}

// StatsManager handles statistics tracking
type StatsManager struct {
	configPath string
	history    map[string][]StatEntry
}

// StatEntry represents a single statistics entry
type StatEntry struct {
	Date      time.Time     `json:"date"`
	WordCount int           `json:"word_count"`
	NoteCount int           `json:"note_count"`
	TimeSaved time.Duration `json:"time_saved"`
}

// NotebookStats represents statistics for a notebook
type NotebookStats struct {
	Name      string
	NoteCount int
	WordCount int
}

// DashboardData holds all dashboard information
type DashboardData struct {
	TotalNotes          int
	TotalWords          int
	TotalNotebooks      int
	CurrentStreak       int
	LongestStreak       int
	WeeklyActivity      map[string]int // day -> note count
	TopNotebooks        []NotebookStats
	MostProductiveDay   string
	AverageWordsPerNote int
}

// NewStatsManager creates a new statistics manager
func NewStatsManager() *StatsManager {
	return &StatsManager{
		history: make(map[string][]StatEntry),
	}
}

// NewStatsManagerWithConfig creates a stats manager with config directory
func NewStatsManagerWithConfig(configDir string) *StatsManager {
	sm := &StatsManager{
		configPath: filepath.Join(configDir, ".stats.json"),
		history:    make(map[string][]StatEntry),
	}
	sm.loadHistory()
	return sm
}

// loadHistory loads statistics history from disk
func (sm *StatsManager) loadHistory() error {
	if sm.configPath == "" {
		return nil
	}

	data, err := os.ReadFile(sm.configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File doesn't exist yet, that's ok
		}
		return err
	}

	return json.Unmarshal(data, &sm.history)
}

// saveHistory saves statistics history to disk
func (sm *StatsManager) saveHistory() error {
	if sm.configPath == "" {
		return nil
	}

	data, err := json.MarshalIndent(sm.history, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(sm.configPath, data, 0644)
}

// Calculate calculates statistics for content
func (sm *StatsManager) Calculate(content string) Statistics {
	stats := Statistics{}

	// Character count
	stats.CharCount = len(content)

	// Line count
	stats.LineCount = len(strings.Split(content, "\n"))

	// Word count
	stats.WordCount = countWords(content)

	// Paragraph count
	stats.ParagraphCount = countParagraphs(content)

	// Sentence count
	stats.SentenceCount = countSentences(content)

	// Reading time (average 200 words per minute)
	stats.ReadingTime = stats.WordCount / 200
	if stats.ReadingTime == 0 && stats.WordCount > 0 {
		stats.ReadingTime = 1
	}

	return stats
}

// RecordStats records statistics for a date
func (sm *StatsManager) RecordStats(date time.Time, wordCount, noteCount int) {
	dateStr := date.Format("2006-01-02")
	entry := StatEntry{
		Date:      date,
		WordCount: wordCount,
		NoteCount: noteCount,
	}

	sm.history[dateStr] = append(sm.history[dateStr], entry)
	sm.saveHistory()
}

// GetStreak calculates the current writing streak
func (sm *StatsManager) GetStreak() int {
	streak := 0
	date := time.Now()

	for {
		dateStr := date.Format("2006-01-02")
		if entries, ok := sm.history[dateStr]; ok && len(entries) > 0 {
			streak++
			date = date.AddDate(0, 0, -1)
		} else {
			break
		}
	}

	return streak
}

// GetTotalWords returns total words written
func (sm *StatsManager) GetTotalWords() int {
	total := 0
	for _, entries := range sm.history {
		for _, entry := range entries {
			total += entry.WordCount
		}
	}
	return total
}

// GetTotalNotes returns total notes created
func (sm *StatsManager) GetTotalNotes() int {
	total := 0
	for _, entries := range sm.history {
		for _, entry := range entries {
			total += entry.NoteCount
		}
	}
	return total
}

// GetLongestStreak calculates the longest writing streak
func (sm *StatsManager) GetLongestStreak() int {
	if len(sm.history) == 0 {
		return 0
	}

	// Get all dates and sort them
	dates := make([]time.Time, 0, len(sm.history))
	for dateStr := range sm.history {
		t, err := time.Parse("2006-01-02", dateStr)
		if err == nil {
			dates = append(dates, t)
		}
	}

	if len(dates) == 0 {
		return 0
	}

	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})

	maxStreak := 1
	currentStreak := 1

	for i := 1; i < len(dates); i++ {
		daysDiff := int(dates[i].Sub(dates[i-1]).Hours() / 24)
		if daysDiff == 1 {
			currentStreak++
			if currentStreak > maxStreak {
				maxStreak = currentStreak
			}
		} else {
			currentStreak = 1
		}
	}

	return maxStreak
}

// GetWeeklyActivity returns note counts for the past 7 days
func (sm *StatsManager) GetWeeklyActivity() map[string]int {
	activity := make(map[string]int)
	now := time.Now()

	for i := 6; i >= 0; i-- {
		date := now.AddDate(0, 0, -i)
		dateStr := date.Format("2006-01-02")
		dayName := date.Format("Mon")

		count := 0
		if entries, ok := sm.history[dateStr]; ok {
			for _, entry := range entries {
				count += entry.NoteCount
			}
		}
		activity[dayName] = count
	}

	return activity
}

// GetMostProductiveDay returns the day with most notes
func (sm *StatsManager) GetMostProductiveDay() string {
	dayCount := make(map[string]int)

	for dateStr, entries := range sm.history {
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			continue
		}

		dayName := t.Format("Monday")
		for _, entry := range entries {
			dayCount[dayName] += entry.NoteCount
		}
	}

	maxDay := "N/A"
	maxCount := 0
	for day, count := range dayCount {
		if count > maxCount {
			maxCount = count
			maxDay = day
		}
	}

	return maxDay
}

// BuildDashboard creates comprehensive dashboard data
func (sm *StatsManager) BuildDashboard(totalNotes, totalNotebooks int, notebookStats []NotebookStats) DashboardData {
	totalWords := sm.GetTotalWords()
	avgWords := 0
	if totalNotes > 0 {
		avgWords = totalWords / totalNotes
	}

	return DashboardData{
		TotalNotes:          totalNotes,
		TotalWords:          totalWords,
		TotalNotebooks:      totalNotebooks,
		CurrentStreak:       sm.GetStreak(),
		LongestStreak:       sm.GetLongestStreak(),
		WeeklyActivity:      sm.GetWeeklyActivity(),
		TopNotebooks:        notebookStats,
		MostProductiveDay:   sm.GetMostProductiveDay(),
		AverageWordsPerNote: avgWords,
	}
}

// RenderDashboard creates a visual representation of dashboard data
func RenderDashboard(data DashboardData) string {
	var sb strings.Builder

	sb.WriteString("📊 WRITING ANALYTICS\n\n")

	// Overall stats
	sb.WriteString("📈 Overview\n")
	sb.WriteString(fmt.Sprintf("  Total Notes: %d\n", data.TotalNotes))
	sb.WriteString(fmt.Sprintf("  Total Words: %s\n", formatNumber(data.TotalWords)))
	sb.WriteString(fmt.Sprintf("  Notebooks: %d\n", data.TotalNotebooks))
	sb.WriteString(fmt.Sprintf("  Avg Words/Note: %d\n\n", data.AverageWordsPerNote))

	// Streaks
	sb.WriteString("🔥 Writing Streaks\n")
	sb.WriteString(fmt.Sprintf("  Current: %d days\n", data.CurrentStreak))
	sb.WriteString(fmt.Sprintf("  Longest: %d days\n", data.LongestStreak))
	sb.WriteString(fmt.Sprintf("  Most Productive: %s\n\n", data.MostProductiveDay))

	// Weekly activity
	if len(data.WeeklyActivity) > 0 {
		sb.WriteString("📅 This Week\n")
		days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
		maxCount := 0
		for _, count := range data.WeeklyActivity {
			if count > maxCount {
				maxCount = count
			}
		}

		for _, day := range days {
			count := data.WeeklyActivity[day]
			bar := createBar(count, maxCount, 15)
			sb.WriteString(fmt.Sprintf("  %s %s %d notes\n", day, bar, count))
		}
		sb.WriteString("\n")
	}

	// Top notebooks
	if len(data.TopNotebooks) > 0 {
		sb.WriteString("🏆 Top Notebooks\n")
		for i, nb := range data.TopNotebooks {
			if i >= 5 {
				break
			}
			sb.WriteString(fmt.Sprintf("  %d. 📁 %s - %d notes\n", i+1, nb.Name, nb.NoteCount))
		}
	}

	return sb.String()
}

// createBar creates a simple ASCII bar chart
func createBar(value, max, width int) string {
	if max == 0 {
		return strings.Repeat("▁", width)
	}
	barWidth := (value * width) / max
	if barWidth == 0 && value > 0 {
		barWidth = 1
	}
	return strings.Repeat("█", barWidth) + strings.Repeat("▁", width-barWidth)
}

// formatNumber formats large numbers with commas
func formatNumber(n int) string {
	if n < 1000 {
		return fmt.Sprintf("%d", n)
	}
	if n < 1000000 {
		return fmt.Sprintf("%d,%03d", n/1000, n%1000)
	}
	return fmt.Sprintf("%d,%03d,%03d", n/1000000, (n%1000000)/1000, n%1000)
}

func countWords(text string) int {
	words := 0
	inWord := false

	for _, r := range text {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			if inWord {
				words++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	if inWord {
		words++
	}

	return words
}

func countParagraphs(text string) int {
	paragraphs := 0
	lines := strings.Split(text, "\n")
	inParagraph := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			if !inParagraph {
				paragraphs++
				inParagraph = true
			}
		} else {
			inParagraph = false
		}
	}

	return paragraphs
}

func countSentences(text string) int {
	count := 0
	for _, r := range text {
		if r == '.' || r == '!' || r == '?' {
			count++
		}
	}
	return count
}
