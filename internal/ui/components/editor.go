package components

import (
	"github.com/charmbracelet/bubbles/textarea"
)

// NewEditor creates a new textarea for note editing
func NewEditor() textarea.Model {
	ta := textarea.New()
	ta.Placeholder = "Write your note here..."
	ta.ShowLineNumbers = false
	ta.Focus()
	return ta
}
