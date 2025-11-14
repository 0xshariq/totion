package components

import (
	"github.com/0xshariq/totion/internal/models"
	"github.com/0xshariq/totion/internal/ui/styles"
	"github.com/charmbracelet/bubbles/list"
)

// NewNoteList creates a new list for displaying notes
func NewNoteList(notes []models.Note) list.Model {
	items := make([]list.Item, len(notes))
	for i, note := range notes {
		items[i] = note
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "All Notes 📒"
	l.Styles.Title = styles.ListTitleStyle

	return l
}
