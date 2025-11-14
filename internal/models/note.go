package models

import (
	"time"
)

// FileFormat represents the file format type
type FileFormat string

const (
	FormatMarkdown FileFormat = "md"
	FormatText     FileFormat = "txt"
)

// Note represents a single note file
type Note struct {
	Name    string
	Path    string
	Format  FileFormat
	Content string
	ModTime time.Time
	Size    int64
}

// Title returns the display title for the note
func (n Note) Title() string {
	return n.Name
}

// Description returns the display description for the note
func (n Note) Description() string {
	formatIcon := "📝"
	if n.Format == FormatText {
		formatIcon = "📄"
	}
	return formatIcon + " " + n.ModTime.Format("2006-01-02 15:04")
}

// FilterValue returns the value used for filtering
func (n Note) FilterValue() string {
	return n.Name
}

// GetExtension returns the file extension for the format
func (f FileFormat) GetExtension() string {
	return "." + string(f)
}

// GetIcon returns the icon for the format
func (f FileFormat) GetIcon() string {
	if f == FormatMarkdown {
		return "📝"
	}
	return "📄"
}

// IsValidFormat checks if the format is valid
func IsValidFormat(format string) bool {
	return format == string(FormatMarkdown) || format == string(FormatText)
}
