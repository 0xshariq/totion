package recent

import (
"encoding/json"
"os"
"path/filepath"
"time"

"github.com/0xshariq/totion/internal/models"
)

type RecentNote struct {
Path     string `json:"path"`
Name     string `json:"name"`
Format   string `json:"format"`
OpenedAt time.Time `json:"opened_at"`
}

type RecentManager struct {
configPath string
maxRecent  int
}

func NewRecentManager(configDir string) *RecentManager {
return &RecentManager{
configPath: filepath.Join(configDir, ".recent_notes.json"),
maxRecent:  10,
}
}

func (r *RecentManager) AddRecent(note *models.Note) error {
recent := r.GetRecent()
for i, n := range recent {
if n.Path == note.Path {
recent = append(recent[:i], recent[i+1:]...)
break
}
}
newRecent := RecentNote{
Path:     note.Path,
Name:     note.Name,
Format:   string(note.Format),
OpenedAt: time.Now(),
}
recent = append([]RecentNote{newRecent}, recent...)
if len(recent) > r.maxRecent {
recent = recent[:r.maxRecent]
}
data, err := json.MarshalIndent(recent, "", "  ")
if err != nil {
return err
}
return os.WriteFile(r.configPath, data, 0644)
}

func (r *RecentManager) GetRecent() []RecentNote {
data, err := os.ReadFile(r.configPath)
if err != nil {
return []RecentNote{}
}
var recent []RecentNote
if err := json.Unmarshal(data, &recent); err != nil {
return []RecentNote{}
}
return recent
}

func (r *RecentManager) Clear() error {
return os.Remove(r.configPath)
}
