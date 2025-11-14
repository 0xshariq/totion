package importpkg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
