package search

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SearchResult represents a search result
type SearchResult struct {
	NotePath     string
	NoteName     string
	LineNumber   int
	MatchSnippet string
	FullLine     string
}

// SearchManager handles full-text search
type SearchManager struct {
	vaultDir string
}

// NewSearchManager creates a new search manager
func NewSearchManager(vaultDir string) *SearchManager {
	return &SearchManager{
		vaultDir: vaultDir,
	}
}

// Search performs full-text search across all notes
func (sm *SearchManager) Search(query string) ([]SearchResult, error) {
	if query == "" {
		return []SearchResult{}, nil
	}

	results := []SearchResult{}
	query = strings.ToLower(query)

	// Walk through all files in vault
	err := filepath.Walk(sm.vaultDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip errors
		}

		// Only search .md and .txt files
		if info.IsDir() || (filepath.Ext(path) != ".md" && filepath.Ext(path) != ".txt") {
			return nil
		}

		// Read file content
		content, err := os.ReadFile(path)
		if err != nil {
			return nil // Skip if can't read
		}

		// Search line by line
		lines := strings.Split(string(content), "\n")
		for lineNum, line := range lines {
			if strings.Contains(strings.ToLower(line), query) {
				snippet := sm.createSnippet(line, query, 50)
				results = append(results, SearchResult{
					NotePath:     path,
					NoteName:     filepath.Base(path),
					LineNumber:   lineNum + 1,
					MatchSnippet: snippet,
					FullLine:     line,
				})

				// Limit results per file to avoid overwhelming
				if len(results) > 100 {
					return filepath.SkipAll
				}
			}
		}

		return nil
	})

	return results, err
}

// createSnippet creates a context snippet around the match
func (sm *SearchManager) createSnippet(line, query string, maxLen int) string {
	lowerLine := strings.ToLower(line)
	lowerQuery := strings.ToLower(query)

	idx := strings.Index(lowerLine, lowerQuery)
	if idx == -1 {
		// Fallback: return beginning of line
		if len(line) > maxLen {
			return line[:maxLen] + "..."
		}
		return line
	}

	// Calculate snippet range
	start := idx - 20
	if start < 0 {
		start = 0
	}

	end := idx + len(query) + 30
	if end > len(line) {
		end = len(line)
	}

	snippet := line[start:end]

	// Add ellipsis if truncated
	if start > 0 {
		snippet = "..." + snippet
	}
	if end < len(line) {
		snippet = snippet + "..."
	}

	return snippet
}

// SearchInNote searches within a specific note
func (sm *SearchManager) SearchInNote(notePath, query string) ([]SearchResult, error) {
	results := []SearchResult{}
	query = strings.ToLower(query)

	content, err := os.ReadFile(notePath)
	if err != nil {
		return results, err
	}

	lines := strings.Split(string(content), "\n")
	for lineNum, line := range lines {
		if strings.Contains(strings.ToLower(line), query) {
			snippet := sm.createSnippet(line, query, 80)
			results = append(results, SearchResult{
				NotePath:     notePath,
				NoteName:     filepath.Base(notePath),
				LineNumber:   lineNum + 1,
				MatchSnippet: snippet,
				FullLine:     line,
			})
		}
	}

	return results, nil
}

// FormatResults formats search results for display
func FormatResults(results []SearchResult) string {
	if len(results) == 0 {
		return "No results found."
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Found %d matches:\n\n", len(results)))

	currentFile := ""
	for i, result := range results {
		if i >= 20 {
			sb.WriteString(fmt.Sprintf("\n... and %d more matches", len(results)-20))
			break
		}

		// Group by file
		if result.NoteName != currentFile {
			if currentFile != "" {
				sb.WriteString("\n")
			}
			sb.WriteString(fmt.Sprintf("📄 %s\n", result.NoteName))
			currentFile = result.NoteName
		}

		sb.WriteString(fmt.Sprintf("  Line %d: %s\n", result.LineNumber, result.MatchSnippet))
	}

	return sb.String()
}

// Count returns the total number of results
func Count(results []SearchResult) int {
	return len(results)
}
