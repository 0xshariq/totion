package notebook

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/0xshariq/totion/internal/features/export"
	"github.com/0xshariq/totion/internal/features/search"
	"github.com/0xshariq/totion/internal/features/tags"
)

// Notebook represents a folder/notebook containing notes
type Notebook struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Parent      *Notebook   `json:"-"`
	Children    []*Notebook `json:"children,omitempty"`
	Icon        string      `json:"icon"`
	NoteCount   int         `json:"note_count"`
	CreatedAt   time.Time   `json:"created_at"`
	ModifiedAt  time.Time   `json:"modified_at"`
	Description string      `json:"description,omitempty"`
	Color       string      `json:"color,omitempty"`
}

// NotebookMetadata stores notebook configuration
type NotebookMetadata struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	Icon        string    `json:"icon"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Tags        []string  `json:"tags,omitempty"`
}

// NotebookManager handles notebook/folder operations
type NotebookManager struct {
	vaultDir string
	root     *Notebook
}

// NewNotebookManager creates a new notebook manager
func NewNotebookManager(vaultDir string) *NotebookManager {
	return &NotebookManager{
		vaultDir: vaultDir,
		root: &Notebook{
			Name: "Root",
			Path: vaultDir,
			Icon: "📁",
		},
	}
}

// CreateNotebook creates a new notebook/folder in ~/.totion/
func (nm *NotebookManager) CreateNotebook(name string) error {
	// Create notebook directory in vault
	notebookPath := filepath.Join(nm.vaultDir, name)

	// Check if notebook already exists
	if _, err := os.Stat(notebookPath); err == nil {
		return fmt.Errorf("notebook '%s' already exists", name)
	}

	// Create directory
	if err := os.MkdirAll(notebookPath, 0755); err != nil {
		return fmt.Errorf("error creating notebook: %w", err)
	}

	// Create metadata file
	metadata := NotebookMetadata{
		Name:       name,
		Icon:       "📓",
		Color:      "blue",
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	if err := nm.saveMetadata(notebookPath, metadata); err != nil {
		return fmt.Errorf("error saving metadata: %w", err)
	}

	return nil
}

// CreateNotebookWithDescription creates a notebook with description
func (nm *NotebookManager) CreateNotebookWithDescription(name, description string) error {
	notebookPath := filepath.Join(nm.vaultDir, name)

	if _, err := os.Stat(notebookPath); err == nil {
		return fmt.Errorf("notebook '%s' already exists", name)
	}

	if err := os.MkdirAll(notebookPath, 0755); err != nil {
		return fmt.Errorf("error creating notebook: %w", err)
	}

	metadata := NotebookMetadata{
		Name:        name,
		Description: description,
		Icon:        "📓",
		Color:       "blue",
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}

	if err := nm.saveMetadata(notebookPath, metadata); err != nil {
		return fmt.Errorf("error saving metadata: %w", err)
	}

	return nil
}

// saveMetadata saves notebook metadata to .notebook.json
func (nm *NotebookManager) saveMetadata(notebookPath string, metadata NotebookMetadata) error {
	metadataPath := filepath.Join(notebookPath, ".notebook.json")
	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(metadataPath, data, 0644)
}

// loadMetadata loads notebook metadata from .notebook.json
func (nm *NotebookManager) loadMetadata(notebookPath string) (*NotebookMetadata, error) {
	metadataPath := filepath.Join(notebookPath, ".notebook.json")
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		// Return default metadata if file doesn't exist
		return &NotebookMetadata{
			Name:       filepath.Base(notebookPath),
			Icon:       "📓",
			Color:      "blue",
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
		}, nil
	}

	var metadata NotebookMetadata
	if err := json.Unmarshal(data, &metadata); err != nil {
		return nil, err
	}

	return &metadata, nil
}

// DeleteNotebook deletes a notebook/folder and all its contents
func (nm *NotebookManager) DeleteNotebook(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("notebook does not exist: %s", path)
	}

	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("error deleting notebook: %w", err)
	}

	return nil
}

// GetNoteCount counts the number of notes in a notebook
func (nm *NotebookManager) GetNoteCount(notebookPath string) (int, error) {
	count := 0
	err := filepath.Walk(notebookPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(path)
			if ext == ".md" || ext == ".txt" {
				count++
			}
		}
		return nil
	})
	return count, err
}

// GetNotebookInfo loads metadata and counts notes
func (nm *NotebookManager) GetNotebookInfo(notebookPath string) (*Notebook, error) {
	metadata, err := nm.loadMetadata(notebookPath)
	if err != nil {
		return nil, err
	}

	noteCount, err := nm.GetNoteCount(notebookPath)
	if err != nil {
		return nil, err
	}

	notebook := &Notebook{
		Name:        metadata.Name,
		Path:        notebookPath,
		Icon:        metadata.Icon,
		NoteCount:   noteCount,
		CreatedAt:   metadata.CreatedAt,
		ModifiedAt:  metadata.ModifiedAt,
		Description: metadata.Description,
		Color:       metadata.Color,
	}

	return notebook, nil
}

// UpdateMetadata updates notebook metadata
func (nm *NotebookManager) UpdateMetadata(notebookPath string, updates map[string]interface{}) error {
	metadata, err := nm.loadMetadata(notebookPath)
	if err != nil {
		return err
	}

	// Update fields
	if name, ok := updates["name"].(string); ok {
		metadata.Name = name
	}
	if description, ok := updates["description"].(string); ok {
		metadata.Description = description
	}
	if icon, ok := updates["icon"].(string); ok {
		metadata.Icon = icon
	}
	if color, ok := updates["color"].(string); ok {
		metadata.Color = color
	}
	if tags, ok := updates["tags"].([]string); ok {
		metadata.Tags = tags
	}

	metadata.ModifiedAt = time.Now()

	return nm.saveMetadata(notebookPath, *metadata)
}

// ListNotebooks lists all notebooks with metadata
func (nm *NotebookManager) ListNotebooks() ([]*Notebook, error) {
	notebooks := []*Notebook{}

	entries, err := os.ReadDir(nm.vaultDir)
	if err != nil {
		return nil, fmt.Errorf("error reading vault directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			notebookPath := filepath.Join(nm.vaultDir, entry.Name())
			notebook, err := nm.GetNotebookInfo(notebookPath)
			if err != nil {
				// If we can't get info, create basic notebook entry
				notebook = &Notebook{
					Name: entry.Name(),
					Path: notebookPath,
					Icon: "📁",
				}
			}
			notebooks = append(notebooks, notebook)
		}
	}

	return notebooks, nil
}

// RenameNotebook renames a notebook and updates its metadata
func (nm *NotebookManager) RenameNotebook(oldPath, newName string) error {
	newPath := filepath.Join(filepath.Dir(oldPath), newName)

	if err := os.Rename(oldPath, newPath); err != nil {
		return fmt.Errorf("error renaming notebook: %w", err)
	}

	// Update metadata with new name
	updates := map[string]interface{}{
		"name": newName,
	}
	if err := nm.UpdateMetadata(newPath, updates); err != nil {
		// Non-fatal error, just log it
		return fmt.Errorf("notebook renamed but metadata update failed: %w", err)
	}

	return nil
}

// MoveNotebook moves a notebook to a new parent
func (nm *NotebookManager) MoveNotebook(notebookPath, newParentPath string) error {
	notebookName := filepath.Base(notebookPath)
	newPath := filepath.Join(newParentPath, notebookName)

	if err := os.Rename(notebookPath, newPath); err != nil {
		return fmt.Errorf("error moving notebook: %w", err)
	}

	return nil
}

// GetNotebookPath returns the full path for a notebook
func (nm *NotebookManager) GetNotebookPath(notebookName string) string {
	return filepath.Join(nm.vaultDir, notebookName)
}

// IsNotebook checks if a path is a notebook/folder
func (nm *NotebookManager) IsNotebook(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}

// GetNotesInNotebook returns all note files in a notebook
func (nm *NotebookManager) GetNotesInNotebook(notebookPath string) ([]string, error) {
	notes := []string{}

	entries, err := os.ReadDir(notebookPath)
	if err != nil {
		return nil, fmt.Errorf("error reading notebook directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			ext := filepath.Ext(entry.Name())
			if ext == ".md" || ext == ".txt" {
				notes = append(notes, filepath.Join(notebookPath, entry.Name()))
			}
		}
	}

	return notes, nil
}

// AddTagToNotebook adds a tag to notebook metadata
func (nm *NotebookManager) AddTagToNotebook(notebookPath string, tag string) error {
	metadata, err := nm.loadMetadata(notebookPath)
	if err != nil {
		return err
	}

	// Check if tag already exists
	for _, existingTag := range metadata.Tags {
		if existingTag == tag {
			return nil // Tag already exists
		}
	}

	metadata.Tags = append(metadata.Tags, tag)
	metadata.ModifiedAt = time.Now()

	return nm.saveMetadata(notebookPath, *metadata)
}

// RemoveTagFromNotebook removes a tag from notebook metadata
func (nm *NotebookManager) RemoveTagFromNotebook(notebookPath string, tag string) error {
	metadata, err := nm.loadMetadata(notebookPath)
	if err != nil {
		return err
	}

	// Find and remove tag
	for i, existingTag := range metadata.Tags {
		if existingTag == tag {
			metadata.Tags = append(metadata.Tags[:i], metadata.Tags[i+1:]...)
			break
		}
	}

	metadata.ModifiedAt = time.Now()

	return nm.saveMetadata(notebookPath, *metadata)
}

// GetNotebooksByTag returns all notebooks with a specific tag
func (nm *NotebookManager) GetNotebooksByTag(tag string) ([]*Notebook, error) {
	allNotebooks, err := nm.ListNotebooks()
	if err != nil {
		return nil, err
	}

	filtered := []*Notebook{}
	for _, notebook := range allNotebooks {
		metadata, err := nm.loadMetadata(notebook.Path)
		if err != nil {
			continue
		}

		for _, t := range metadata.Tags {
			if t == tag {
				filtered = append(filtered, notebook)
				break
			}
		}
	}

	return filtered, nil
}

// SearchInNotebook searches for text within a notebook's notes
func (nm *NotebookManager) SearchInNotebook(notebookPath, query string) ([]search.SearchResult, error) {
	searchMgr := search.NewSearchManager(notebookPath)
	return searchMgr.Search(query)
}

// GetTagsInNotebook gets all unique tags from notes in a notebook
func (nm *NotebookManager) GetTagsInNotebook(notebookPath string) ([]string, error) {
	tagSet := make(map[string]bool)

	notes, err := nm.GetNotesInNotebook(notebookPath)
	if err != nil {
		return nil, err
	}

	for _, notePath := range notes {
		content, err := os.ReadFile(notePath)
		if err != nil {
			continue
		}

		noteTags := tags.ExtractTags(string(content))
		for _, tag := range noteTags {
			tagSet[tag] = true
		}
	}

	// Convert to slice
	uniqueTags := make([]string, 0, len(tagSet))
	for tag := range tagSet {
		uniqueTags = append(uniqueTags, tag)
	}

	return uniqueTags, nil
}

// ExportNotebook exports all notes in a notebook to a specified format
func (nm *NotebookManager) ExportNotebook(notebookPath, outputDir string, format export.ExportFormat) error {
	notes, err := nm.GetNotesInNotebook(notebookPath)
	if err != nil {
		return fmt.Errorf("error getting notes: %w", err)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	exporter := export.NewExporter()
	exportData := []export.NoteData{}

	for _, notePath := range notes {
		content, err := os.ReadFile(notePath)
		if err != nil {
			continue
		}

		title := strings.TrimSuffix(filepath.Base(notePath), filepath.Ext(notePath))
		exportData = append(exportData, export.NoteData{
			Title:   title,
			Content: string(content),
			Path:    notePath,
		})
	}

	return exporter.BatchExport(exportData, outputDir, format)
}

// GetNotebookStatistics returns statistics for a notebook
func (nm *NotebookManager) GetNotebookStatistics(notebookPath string) (*NotebookStats, error) {
	notes, err := nm.GetNotesInNotebook(notebookPath)
	if err != nil {
		return nil, err
	}

	stats := &NotebookStats{
		TotalNotes: len(notes),
	}

	for _, notePath := range notes {
		info, err := os.Stat(notePath)
		if err != nil {
			continue
		}

		stats.TotalSize += info.Size()

		content, err := os.ReadFile(notePath)
		if err != nil {
			continue
		}

		words := len(strings.Fields(string(content)))
		stats.TotalWords += words
	}

	tags, _ := nm.GetTagsInNotebook(notebookPath)
	stats.UniqueTags = len(tags)

	return stats, nil
}

// NotebookStats represents notebook statistics
type NotebookStats struct {
	TotalNotes int
	TotalSize  int64
	TotalWords int
	UniqueTags int
}

// DuplicateNotebook creates a copy of a notebook
func (nm *NotebookManager) DuplicateNotebook(srcPath, newName string) error {
	destPath := filepath.Join(filepath.Dir(srcPath), newName)

	if _, err := os.Stat(destPath); err == nil {
		return fmt.Errorf("notebook '%s' already exists", newName)
	}

	// Copy directory recursively
	return nm.copyDir(srcPath, destPath)
}

// copyDir recursively copies a directory
func (nm *NotebookManager) copyDir(src, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := nm.copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := nm.copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// copyFile copies a single file
func (nm *NotebookManager) copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, data, info.Mode())
}

// ArchiveNotebook creates a zip archive of a notebook
func (nm *NotebookManager) ArchiveNotebook(notebookPath, archivePath string) error {
	// Simple archive by exporting all notes
	return nm.ExportNotebook(notebookPath, archivePath, export.FormatMarkdown)
}

// GetRecentlyModifiedNotes returns recently modified notes in a notebook
func (nm *NotebookManager) GetRecentlyModifiedNotes(notebookPath string, limit int) ([]string, error) {
	type NoteInfo struct {
		Path    string
		ModTime time.Time
	}

	notes, err := nm.GetNotesInNotebook(notebookPath)
	if err != nil {
		return nil, err
	}

	noteInfos := []NoteInfo{}
	for _, notePath := range notes {
		info, err := os.Stat(notePath)
		if err != nil {
			continue
		}

		noteInfos = append(noteInfos, NoteInfo{
			Path:    notePath,
			ModTime: info.ModTime(),
		})
	}

	// Sort by modification time (newest first)
	for i := 0; i < len(noteInfos)-1; i++ {
		for j := i + 1; j < len(noteInfos); j++ {
			if noteInfos[i].ModTime.Before(noteInfos[j].ModTime) {
				noteInfos[i], noteInfos[j] = noteInfos[j], noteInfos[i]
			}
		}
	}

	// Return top N
	result := []string{}
	for i := 0; i < len(noteInfos) && i < limit; i++ {
		result = append(result, noteInfos[i].Path)
	}

	return result, nil
}
