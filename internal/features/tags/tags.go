package tags

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// TagInfo represents a tag with its associated notes
type TagInfo struct {
	Tag   string   `json:"tag"`
	Notes []string `json:"notes"` // List of note paths
	Count int      `json:"count"`
}

// TagManager handles tag extraction and indexing
type TagManager struct {
	vaultDir  string
	indexPath string
	tags      map[string]*TagInfo // tag -> TagInfo
}

// NewTagManager creates a new tag manager
func NewTagManager(vaultDir, configDir string) *TagManager {
	tm := &TagManager{
		vaultDir:  vaultDir,
		indexPath: filepath.Join(configDir, ".tags.json"),
		tags:      make(map[string]*TagInfo),
	}
	tm.loadIndex()
	return tm
}

// ExtractTags extracts all #tags from content
func ExtractTags(content string) []string {
	// Match #word (but not ##heading or #123)
	tagPattern := regexp.MustCompile(`(?:^|[^\w#])#([a-zA-Z][a-zA-Z0-9_-]*)\b`)
	matches := tagPattern.FindAllStringSubmatch(content, -1)
	
	tags := []string{}
	seen := make(map[string]bool)
	
	for _, match := range matches {
		if len(match) > 1 {
			tag := strings.ToLower(match[1])
			if !seen[tag] {
				tags = append(tags, tag)
				seen[tag] = true
			}
		}
	}
	
	return tags
}

// RebuildIndex rebuilds the entire tag index by scanning all notes
func (tm *TagManager) RebuildIndex() error {
	tm.tags = make(map[string]*TagInfo)
	
	err := filepath.Walk(tm.vaultDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		
		// Only index .md and .txt files
		if info.IsDir() || (filepath.Ext(path) != ".md" && filepath.Ext(path) != ".txt") {
			return nil
		}
		
		// Read and extract tags
		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		
		tags := ExtractTags(string(content))
		for _, tag := range tags {
			if _, exists := tm.tags[tag]; !exists {
				tm.tags[tag] = &TagInfo{
					Tag:   tag,
					Notes: []string{},
					Count: 0,
				}
			}
			tm.tags[tag].Notes = append(tm.tags[tag].Notes, path)
			tm.tags[tag].Count++
		}
		
		return nil
	})
	
	if err != nil {
		return err
	}
	
	return tm.saveIndex()
}

// GetNotesByTag returns all notes containing a specific tag
func (tm *TagManager) GetNotesByTag(tag string) []string {
	tag = strings.ToLower(tag)
	if info, exists := tm.tags[tag]; exists {
		return info.Notes
	}
	return []string{}
}

// GetAllTags returns all tags sorted by frequency
func (tm *TagManager) GetAllTags() []*TagInfo {
	tags := make([]*TagInfo, 0, len(tm.tags))
	for _, info := range tm.tags {
		tags = append(tags, info)
	}
	
	// Sort by count (descending)
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Count > tags[j].Count
	})
	
	return tags
}

// GetTagsForNote returns all tags in a specific note
func (tm *TagManager) GetTagsForNote(notePath string) ([]string, error) {
	content, err := os.ReadFile(notePath)
	if err != nil {
		return nil, err
	}
	
	return ExtractTags(string(content)), nil
}

// IndexNote updates the index for a specific note
func (tm *TagManager) IndexNote(notePath string) error {
	// Remove this note from all existing tags
	for _, info := range tm.tags {
		newNotes := []string{}
		for _, note := range info.Notes {
			if note != notePath {
				newNotes = append(newNotes, note)
			}
		}
		info.Notes = newNotes
		info.Count = len(newNotes)
	}
	
	// Re-index the note
	content, err := os.ReadFile(notePath)
	if err != nil {
		return err
	}
	
	tags := ExtractTags(string(content))
	for _, tag := range tags {
		if _, exists := tm.tags[tag]; !exists {
			tm.tags[tag] = &TagInfo{
				Tag:   tag,
				Notes: []string{},
				Count: 0,
			}
		}
		tm.tags[tag].Notes = append(tm.tags[tag].Notes, notePath)
		tm.tags[tag].Count++
	}
	
	return tm.saveIndex()
}

// FormatTagCloud formats tags for display
func FormatTagCloud(tags []*TagInfo, maxTags int) string {
	if len(tags) == 0 {
		return "No tags found."
	}
	
	var sb strings.Builder
	sb.WriteString("📑 Tag Cloud:\n\n")
	
	count := 0
	for _, tag := range tags {
		if count >= maxTags {
			break
		}
		
		// Create visual representation of frequency
		bars := ""
		numBars := tag.Count
		if numBars > 10 {
			numBars = 10
		}
		for i := 0; i < numBars; i++ {
			bars += "▪"
		}
		
		sb.WriteString(fmt.Sprintf("  #%-20s %s (%d)\n", tag.Tag, bars, tag.Count))
		count++
	}
	
	if len(tags) > maxTags {
		sb.WriteString(fmt.Sprintf("\n... and %d more tags", len(tags)-maxTags))
	}
	
	return sb.String()
}

// FormatTagsForNote formats tags found in a note
func FormatTagsForNote(tags []string) string {
	if len(tags) == 0 {
		return "No tags"
	}
	
	formatted := make([]string, len(tags))
	for i, tag := range tags {
		formatted[i] = "#" + tag
	}
	
	return strings.Join(formatted, " ")
}

// loadIndex loads the tag index from disk
func (tm *TagManager) loadIndex() error {
	data, err := os.ReadFile(tm.indexPath)
	if err != nil {
		// File doesn't exist yet, that's okay
		return nil
	}
	
	var tagList []*TagInfo
	if err := json.Unmarshal(data, &tagList); err != nil {
		return err
	}
	
	// Convert to map
	for _, info := range tagList {
		tm.tags[info.Tag] = info
	}
	
	return nil
}

// saveIndex saves the tag index to disk
func (tm *TagManager) saveIndex() error {
	// Convert map to list
	tagList := make([]*TagInfo, 0, len(tm.tags))
	for _, info := range tm.tags {
		if info.Count > 0 { // Only save tags with notes
			tagList = append(tagList, info)
		}
	}
	
	data, err := json.MarshalIndent(tagList, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(tm.indexPath, data, 0644)
}

// GetTagCount returns the total number of unique tags
func (tm *TagManager) GetTagCount() int {
	count := 0
	for _, info := range tm.tags {
		if info.Count > 0 {
			count++
		}
	}
	return count
}
