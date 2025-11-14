package quick

import (
	"os"
	"path/filepath"
	"time"
)

// QuickNoteManager handles quick/scratch notes
type QuickNoteManager struct {
	scratchPath string
}

// NewQuickNoteManager creates a new quick note manager
func NewQuickNoteManager(configDir string) *QuickNoteManager {
	return &QuickNoteManager{
		scratchPath: filepath.Join(configDir, ".scratch.md"),
	}
}

// GetScratchPath returns the scratch pad file path
func (qm *QuickNoteManager) GetScratchPath() string {
	return qm.scratchPath
}

// LoadScratch loads scratch pad content
func (qm *QuickNoteManager) LoadScratch() (string, error) {
	data, err := os.ReadFile(qm.scratchPath)
	if err != nil {
		if os.IsNotExist(err) {
			return qm.getDefaultContent(), nil
		}
		return "", err
	}
	return string(data), nil
}

// SaveScratch saves scratch pad content
func (qm *QuickNoteManager) SaveScratch(content string) error {
	return os.WriteFile(qm.scratchPath, []byte(content), 0644)
}

// ClearScratch clears the scratch pad
func (qm *QuickNoteManager) ClearScratch() error {
	return qm.SaveScratch(qm.getDefaultContent())
}

// getDefaultContent returns default scratch content
func (qm *QuickNoteManager) getDefaultContent() string {
	return `# 📝 Quick Notes

Jot down quick thoughts, ideas, or temporary notes here.
This is your scratch pad - notes here are auto-saved.

---

`
}

// PromoteToNote promotes scratch content to a real note
func (qm *QuickNoteManager) PromoteToNote(vaultDir, noteName string) (string, error) {
	// Read scratch content
	content, err := qm.LoadScratch()
	if err != nil {
		return "", err
	}

	// Create note file
	notePath := filepath.Join(vaultDir, noteName+".md")
	if err := os.WriteFile(notePath, []byte(content), 0644); err != nil {
		return "", err
	}

	// Clear scratch
	qm.ClearScratch()

	return notePath, nil
}

// AutoSave saves scratch content every interval
func (qm *QuickNoteManager) AutoSave(content string) error {
	return qm.SaveScratch(content)
}

// GetLastModified returns when scratch was last modified
func (qm *QuickNoteManager) GetLastModified() (time.Time, error) {
	info, err := os.Stat(qm.scratchPath)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}
