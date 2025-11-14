package daily

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// DailyManager handles daily notes
type DailyManager struct {
	vaultDir string
}

// NewDailyManager creates a new daily notes manager
func NewDailyManager(vaultDir string) *DailyManager {
	return &DailyManager{
		vaultDir: vaultDir,
	}
}

// GetTodayNotePath returns the path for today's note
func (dm *DailyManager) GetTodayNotePath() string {
	today := time.Now().Format("2006-01-02")
	dailyDir := filepath.Join(dm.vaultDir, "Daily")
	return filepath.Join(dailyDir, today+".md")
}

// CreateTodayNote creates or opens today's daily note
func (dm *DailyManager) CreateTodayNote() (string, error) {
	notePath := dm.GetTodayNotePath()
	dailyDir := filepath.Dir(notePath)

	// Create Daily directory if it doesn't exist
	if err := os.MkdirAll(dailyDir, 0755); err != nil {
		return "", err
	}

	// Check if note already exists
	if _, err := os.Stat(notePath); err == nil {
		// Note already exists
		return notePath, nil
	}

	// Create new daily note with template
	content := dm.getDailyTemplate()
	if err := os.WriteFile(notePath, []byte(content), 0644); err != nil {
		return "", err
	}

	return notePath, nil
}

// getDailyTemplate returns the daily note template
func (dm *DailyManager) getDailyTemplate() string {
	now := time.Now()
	date := now.Format("January 02, 2006")
	day := now.Format("Monday")

	return fmt.Sprintf(`# 📅 %s (%s)

## 🎯 Today's Goals
- [ ] 
- [ ] 
- [ ] 

## 📝 Notes


## 💭 Reflections


## ✅ Completed
- [ ] 

---
*Created: %s*
`, date, day, now.Format("15:04"))
}

// TodayNoteExists checks if today's note exists
func (dm *DailyManager) TodayNoteExists() bool {
	notePath := dm.GetTodayNotePath()
	_, err := os.Stat(notePath)
	return err == nil
}

// GetDailyNotesCount returns the count of daily notes
func (dm *DailyManager) GetDailyNotesCount() int {
	dailyDir := filepath.Join(dm.vaultDir, "Daily")
	entries, err := os.ReadDir(dailyDir)
	if err != nil {
		return 0
	}

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".md" {
			count++
		}
	}
	return count
}
