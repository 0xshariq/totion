package importpkg

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Importer handles importing notes from various sources
type Importer struct {
	vaultDir string
}

// NewImporter creates a new importer
func NewImporter(vaultDir string) *Importer {
	return &Importer{
		vaultDir: vaultDir,
	}
}

// NotionPage represents a Notion page structure
type NotionPage struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ObsidianNote represents an Obsidian note
type ObsidianNote struct {
	Name    string
	Content string
}

// ImportFromNotion imports notes from Notion export
func (i *Importer) ImportFromNotion(jsonPath string) ([]string, error) {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var pages []NotionPage
	if err := json.Unmarshal(data, &pages); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	imported := []string{}
	for _, page := range pages {
		filename := i.sanitizeFilename(page.Title) + ".md"
		path := filepath.Join(i.vaultDir, "imported", filename)

		if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
			continue
		}

		if err := os.WriteFile(path, []byte(page.Content), 0644); err != nil {
			continue
		}

		imported = append(imported, filename)
	}

	return imported, nil
}

// ImportFromObsidian imports notes from Obsidian vault
func (i *Importer) ImportFromObsidian(obsidianVaultPath string) ([]string, error) {
	imported := []string{}

	err := filepath.Walk(obsidianVaultPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil // Skip files we can't read
		}

		destPath := filepath.Join(i.vaultDir, "imported", info.Name())

		if err := os.MkdirAll(filepath.Dir(destPath), 0750); err != nil {
			return nil
		}

		if err := os.WriteFile(destPath, content, 0644); err != nil {
			return nil
		}

		imported = append(imported, info.Name())
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}

	return imported, nil
}

// ImportFromPlainText imports plain text files
func (i *Importer) ImportFromPlainText(txtPath string) error {
	content, err := os.ReadFile(txtPath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	filename := filepath.Base(txtPath)
	destPath := filepath.Join(i.vaultDir, "imported", filename)

	if err := os.MkdirAll(filepath.Dir(destPath), 0750); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	if err := os.WriteFile(destPath, content, 0644); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

// sanitizeFilename removes invalid characters from filename
func (i *Importer) sanitizeFilename(name string) string {
	// Replace invalid characters
	name = strings.ReplaceAll(name, "/", "-")
	name = strings.ReplaceAll(name, "\\", "-")
	name = strings.ReplaceAll(name, ":", "-")
	name = strings.ReplaceAll(name, "*", "-")
	name = strings.ReplaceAll(name, "?", "-")
	name = strings.ReplaceAll(name, "\"", "-")
	name = strings.ReplaceAll(name, "<", "-")
	name = strings.ReplaceAll(name, ">", "-")
	name = strings.ReplaceAll(name, "|", "-")

	return name
}

// ImportFromJSON imports notes from JSON export
func (i *Importer) ImportFromJSON(jsonPath string) ([]string, error) {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var noteData map[string]interface{}
	if err := json.Unmarshal(data, &noteData); err != nil {
		// Try array of notes
		var notesArray []map[string]interface{}
		if err := json.Unmarshal(data, &notesArray); err != nil {
			return nil, fmt.Errorf("error parsing JSON: %w", err)
		}

		imported := []string{}
		for _, note := range notesArray {
			title, _ := note["title"].(string)
			content, _ := note["content"].(string)

			if title == "" || content == "" {
				continue
			}

			filename := i.sanitizeFilename(title) + ".md"
			path := filepath.Join(i.vaultDir, "imported", filename)

			if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
				continue
			}

			if err := os.WriteFile(path, []byte(content), 0644); err != nil {
				continue
			}

			imported = append(imported, filename)
		}
		return imported, nil
	}

	// Single note JSON
	title, _ := noteData["title"].(string)
	content, _ := noteData["content"].(string)

	if title == "" || content == "" {
		return nil, fmt.Errorf("invalid JSON format: missing title or content")
	}

	filename := i.sanitizeFilename(title) + ".md"
	path := filepath.Join(i.vaultDir, "imported", filename)

	if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
		return nil, fmt.Errorf("error creating directory: %w", err)
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return nil, fmt.Errorf("error writing file: %w", err)
	}

	return []string{filename}, nil
}

// ImportFromCSV imports notes from CSV file (title, content columns)
func (i *Importer) ImportFromCSV(csvPath string) ([]string, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %w", err)
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("CSV file is empty or has no data")
	}

	imported := []string{}
	// Skip header row
	for _, record := range records[1:] {
		if len(record) < 2 {
			continue
		}

		title := record[0]
		content := record[1]

		filename := i.sanitizeFilename(title) + ".md"
		path := filepath.Join(i.vaultDir, "imported", filename)

		if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
			continue
		}

		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			continue
		}

		imported = append(imported, filename)
	}

	return imported, nil
}

// BatchImportFromDirectory imports all markdown/text files from a directory
func (i *Importer) BatchImportFromDirectory(sourcePath string) ([]string, error) {
	imported := []string{}

	err := filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(info.Name())
		if ext != ".md" && ext != ".txt" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		// Preserve relative directory structure
		relPath, _ := filepath.Rel(sourcePath, path)
		destPath := filepath.Join(i.vaultDir, "imported", relPath)

		if err := os.MkdirAll(filepath.Dir(destPath), 0750); err != nil {
			return nil
		}

		if err := os.WriteFile(destPath, content, 0644); err != nil {
			return nil
		}

		imported = append(imported, relPath)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}

	return imported, nil
}

// ImportStats represents import statistics
type ImportStats struct {
	TotalFiles    int
	ImportedFiles int
	SkippedFiles  int
	Errors        []string
	Duration      time.Duration
}

// BatchImportWithStats imports files and returns detailed statistics
func (i *Importer) BatchImportWithStats(sourcePath string) (*ImportStats, error) {
	start := time.Now()
	stats := &ImportStats{
		Errors: []string{},
	}

	err := filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(info.Name())
		if ext != ".md" && ext != ".txt" {
			stats.SkippedFiles++
			return nil
		}

		stats.TotalFiles++

		content, err := os.ReadFile(path)
		if err != nil {
			stats.Errors = append(stats.Errors, fmt.Sprintf("%s: %v", info.Name(), err))
			return nil
		}

		relPath, _ := filepath.Rel(sourcePath, path)
		destPath := filepath.Join(i.vaultDir, "imported", relPath)

		if err := os.MkdirAll(filepath.Dir(destPath), 0750); err != nil {
			stats.Errors = append(stats.Errors, fmt.Sprintf("%s: %v", info.Name(), err))
			return nil
		}

		if err := os.WriteFile(destPath, content, 0644); err != nil {
			stats.Errors = append(stats.Errors, fmt.Sprintf("%s: %v", info.Name(), err))
			return nil
		}

		stats.ImportedFiles++
		return nil
	})

	stats.Duration = time.Since(start)

	if err != nil {
		return stats, fmt.Errorf("error walking directory: %w", err)
	}

	return stats, nil
}
