package templates

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Template represents a note template
type Template struct {
	Name    string
	Content string
	Icon    string
}

// TemplateManager handles note templates
type TemplateManager struct {
	templates  []Template
	configPath string
}

// NewTemplateManager creates a new template manager
func NewTemplateManager() *TemplateManager {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".totion", ".custom_templates.json")

	tm := &TemplateManager{
		configPath: configPath,
	}
	tm.initializeDefaultTemplates()
	tm.loadCustomTemplates()
	return tm
}

// initializeDefaultTemplates sets up default templates
func (tm *TemplateManager) initializeDefaultTemplates() {
	tm.templates = []Template{
		{
			Name: "Meeting Notes",
			Icon: "📋",
			Content: fmt.Sprintf(`# Meeting Notes - %s

## Attendees
- 

## Agenda
1. 

## Discussion Points


## Action Items
- [ ] 

## Next Steps


---
Date: %s
`, time.Now().Format("2006-01-02"), time.Now().Format("15:04")),
		},
		{
			Name: "Todo List",
			Icon: "✅",
			Content: `# Todo List

## Today
- [ ] 

## This Week
- [ ] 

## Later
- [ ] 

---
Created: ` + time.Now().Format("2006-01-02 15:04"),
		},
		{
			Name: "Journal Entry",
			Icon: "📔",
			Content: fmt.Sprintf(`# Journal Entry - %s

## Mood
😊 

## Highlights


## Challenges


## Gratitude
- 

## Tomorrow's Goals
- [ ] 

---
`, time.Now().Format("Monday, January 2, 2006")),
		},
		{
			Name: "Project Plan",
			Icon: "🚀",
			Content: `# Project Plan

## Overview


## Goals
1. 

## Timeline


## Resources Needed


## Milestones
- [ ] 

## Risks


---
Created: ` + time.Now().Format("2006-01-02"),
		},
		{
			Name:    "Code Snippet",
			Icon:    "💻",
			Content: "# Code Snippet\n\n## Description\n\n\n## Language\n\n\n## Code\n```\n\n```\n\n## Notes\n\n",
		},
		{
			Name: "Book Notes",
			Icon: "📚",
			Content: `# Book Notes

## Title


## Author


## Key Takeaways
- 

## Quotes


## My Thoughts


## Rating
⭐ /5

---
Date Read: ` + time.Now().Format("2006-01-02"),
		},
		{
			Name:    "Blank",
			Icon:    "📄",
			Content: "",
		},
	}
}

// GetTemplates returns all available templates
func (tm *TemplateManager) GetTemplates() []Template {
	return tm.templates
}

// GetTemplate returns a specific template by name
func (tm *TemplateManager) GetTemplate(name string) (Template, error) {
	for _, t := range tm.templates {
		if t.Name == name {
			return t, nil
		}
	}
	return Template{}, fmt.Errorf("template not found: %s", name)
}

// AddCustomTemplate adds a user-defined template
func (tm *TemplateManager) AddCustomTemplate(name, content, icon string) {
	tm.templates = append(tm.templates, Template{
		Name:    name,
		Content: content,
		Icon:    icon,
	})
}

// SaveCustomTemplate saves a custom template to disk
func (tm *TemplateManager) SaveCustomTemplate(name, content, icon string) error {
	// Add to current session
	tm.AddCustomTemplate(name, content, icon)

	// Load existing custom templates
	customTemplates := tm.loadCustomTemplatesFromFile()

	// Add new template
	customTemplates = append(customTemplates, Template{
		Name:    name,
		Content: content,
		Icon:    icon,
	})

	// Save to file
	data, err := json.MarshalIndent(customTemplates, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling templates: %w", err)
	}

	// Ensure directory exists
	dir := filepath.Dir(tm.configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	if err := os.WriteFile(tm.configPath, data, 0644); err != nil {
		return fmt.Errorf("error writing templates file: %w", err)
	}

	return nil
}

// loadCustomTemplates loads custom templates from disk
func (tm *TemplateManager) loadCustomTemplates() {
	customTemplates := tm.loadCustomTemplatesFromFile()
	tm.templates = append(tm.templates, customTemplates...)
}

// loadCustomTemplatesFromFile reads custom templates from file
func (tm *TemplateManager) loadCustomTemplatesFromFile() []Template {
	data, err := os.ReadFile(tm.configPath)
	if err != nil {
		return []Template{}
	}

	var templates []Template
	if err := json.Unmarshal(data, &templates); err != nil {
		return []Template{}
	}

	return templates
}

// DeleteCustomTemplate deletes a custom template
func (tm *TemplateManager) DeleteCustomTemplate(name string) error {
	customTemplates := tm.loadCustomTemplatesFromFile()

	// Find and remove template
	for i, t := range customTemplates {
		if t.Name == name {
			customTemplates = append(customTemplates[:i], customTemplates[i+1:]...)
			break
		}
	}

	// Save updated list
	data, err := json.MarshalIndent(customTemplates, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling templates: %w", err)
	}

	if err := os.WriteFile(tm.configPath, data, 0644); err != nil {
		return fmt.Errorf("error writing templates file: %w", err)
	}

	// Reload templates
	tm.templates = tm.templates[:7] // Keep only default templates
	tm.loadCustomTemplates()

	return nil
}
