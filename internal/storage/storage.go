package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/0xshariq/totion/internal/config"
	"github.com/0xshariq/totion/internal/models"
)

// Storage handles file operations
type Storage struct {
	vaultDir string
}

// New creates a new Storage instance
func New() *Storage {
	return &Storage{
		vaultDir: config.AppConfig.VaultDir,
	}
}

// ListNotes returns all notes in the vault
func (s *Storage) ListNotes() ([]models.Note, error) {
	entries, err := os.ReadDir(s.vaultDir)
	if err != nil {
		return nil, fmt.Errorf("error reading vault directory: %w", err)
	}

	notes := make([]models.Note, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		// Only process .md and .txt files
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext != ".md" && ext != ".txt" {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		format := models.FormatMarkdown
		if ext == ".txt" {
			format = models.FormatText
		}

		notes = append(notes, models.Note{
			Name:    entry.Name(),
			Path:    filepath.Join(s.vaultDir, entry.Name()),
			Format:  format,
			ModTime: info.ModTime(),
			Size:    info.Size(),
		})
	}

	return notes, nil
}

// CreateNote creates a new note with the specified name and format
func (s *Storage) CreateNote(name string, format models.FileFormat) (*os.File, string, error) {
	filename := name + format.GetExtension()
	filepath := filepath.Join(s.vaultDir, filename)

	// Check if file already exists
	if _, err := os.Stat(filepath); err == nil {
		return nil, "", fmt.Errorf("file already exists: %s", filename)
	}

	f, err := os.Create(filepath)
	if err != nil {
		return nil, "", fmt.Errorf("error creating file: %w", err)
	}

	return f, filepath, nil
}

// ReadNote reads the content of a note
func (s *Storage) ReadNote(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(content), nil
}

// OpenNote opens a note for reading and writing
func (s *Storage) OpenNote(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return f, nil
}

// SaveNote saves the content to the file
func (s *Storage) SaveNote(file *os.File, content string) error {
	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("error truncating file: %w", err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking file: %w", err)
	}

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("error closing file: %w", err)
	}

	return nil
}

// DeleteNote deletes a note by its path
func (s *Storage) DeleteNote(path string) error {
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("error deleting file: %w", err)
	}

	return nil
}

// FilterByFormat filters notes by format
func (s *Storage) FilterByFormat(notes []models.Note, format models.FileFormat) []models.Note {
	filtered := make([]models.Note, 0)
	for _, note := range notes {
		if note.Format == format {
			filtered = append(filtered, note)
		}
	}
	return filtered
}
