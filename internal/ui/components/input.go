package components

import (
	"github.com/0xshariq/totion/internal/ui/styles"
	"github.com/charmbracelet/bubbles/textinput"
)

// NewFileNameInput creates a new text input for file names
func NewFileNameInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Enter note name..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	ti.Cursor.Style = styles.CursorStyle
	ti.PromptStyle = styles.CursorStyle
	ti.TextStyle = styles.CursorStyle
	return ti
}
