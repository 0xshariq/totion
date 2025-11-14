package pinned

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// PinnedNote represents a pinned note
type PinnedNote struct {
	Path     string    `json:"path"`
	Name     string    `json:"name"`
	PinnedAt time.Time `json:"pinned_at"`
}

// PinnedManager manages pinned notes
type PinnedManager struct {
	configPath string
	pinned     []PinnedNote
}

// NewPinnedManager creates a new pinned notes manager
func NewPinnedManager(configDir string) *PinnedManager {
	pm := &PinnedManager{
		configPath: filepath.Join(configDir, ".pinned_notes.json"),
		pinned:     []PinnedNote{},
	}
	pm.load()
	return pm
}

// load loads pinned notes from disk
func (pm *PinnedManager) load() error {
	data, err := os.ReadFile(pm.configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File doesn't exist yet
		}
		return err
	}

	return json.Unmarshal(data, &pm.pinned)
}

// save saves pinned notes to disk
func (pm *PinnedManager) save() error {
	data, err := json.MarshalIndent(pm.pinned, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(pm.configPath, data, 0644)
}

// Pin pins a note
func (pm *PinnedManager) Pin(path, name string) error {
	// Check if already pinned
	for _, p := range pm.pinned {
		if p.Path == path {
			return nil // Already pinned
		}
	}

	// Limit to 10 pinned notes
	if len(pm.pinned) >= 10 {
		// Remove oldest
		pm.pinned = pm.pinned[1:]
	}

	pm.pinned = append(pm.pinned, PinnedNote{
		Path:     path,
		Name:     name,
		PinnedAt: time.Now(),
	})

	return pm.save()
}

// Unpin unpins a note
func (pm *PinnedManager) Unpin(path string) error {
	for i, p := range pm.pinned {
		if p.Path == path {
			pm.pinned = append(pm.pinned[:i], pm.pinned[i+1:]...)
			return pm.save()
		}
	}
	return nil
}

// IsPinned checks if a note is pinned
func (pm *PinnedManager) IsPinned(path string) bool {
	for _, p := range pm.pinned {
		if p.Path == path {
			return true
		}
	}
	return false
}

// GetPinned returns all pinned notes
func (pm *PinnedManager) GetPinned() []PinnedNote {
	return pm.pinned
}

// Toggle toggles pin status of a note
func (pm *PinnedManager) Toggle(path, name string) (bool, error) {
	if pm.IsPinned(path) {
		err := pm.Unpin(path)
		return false, err
	}
	err := pm.Pin(path, name)
	return true, err
}

// Clear removes all pinned notes
func (pm *PinnedManager) Clear() error {
	pm.pinned = []PinnedNote{}
	return pm.save()
}

// Count returns the number of pinned notes
func (pm *PinnedManager) Count() int {
	return len(pm.pinned)
}
