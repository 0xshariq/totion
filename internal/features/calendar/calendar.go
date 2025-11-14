package calendar

import (
	"fmt"
	"path/filepath"
	"time"
)

// CalendarManager handles daily notes and calendar features
type CalendarManager struct {
	vaultDir string
}

// NewCalendarManager creates a new calendar manager
func NewCalendarManager(vaultDir string) *CalendarManager {
	return &CalendarManager{
		vaultDir: vaultDir,
	}
}

// GetDailyNotePath returns the path for today's daily note
func (cm *CalendarManager) GetDailyNotePath() string {
	date := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("daily-%s.md", date)
	return filepath.Join(cm.vaultDir, "daily", filename)
}

// GetDailyNotePathForDate returns the path for a specific date's daily note
func (cm *CalendarManager) GetDailyNotePathForDate(date time.Time) string {
	dateStr := date.Format("2006-01-02")
	filename := fmt.Sprintf("daily-%s.md", dateStr)
	return filepath.Join(cm.vaultDir, "daily", filename)
}

// CreateDailyNote creates a daily note with a template
func (cm *CalendarManager) CreateDailyNote() string {
	now := time.Now()
	template := fmt.Sprintf(`# Daily Note - %s

## Tasks
- [ ] 

## Notes


## Reflections


---
Created: %s
`, now.Format("Monday, January 2, 2006"), now.Format("15:04"))

	return template
}

// GetWeekRange returns the start and end dates for the current week
func (cm *CalendarManager) GetWeekRange() (start, end time.Time) {
	now := time.Now()
	weekday := int(now.Weekday())

	// Start of week (Sunday)
	start = now.AddDate(0, 0, -weekday)
	// End of week (Saturday)
	end = start.AddDate(0, 0, 6)

	return
}

// FormatDateForDisplay formats a date for display
func (cm *CalendarManager) FormatDateForDisplay(date time.Time) string {
	return date.Format("Mon, Jan 2, 2006")
}
