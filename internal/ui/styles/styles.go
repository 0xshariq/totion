package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Simple 3-color palette
	ColorBlue  = lipgloss.Color("39")  // Blue
	ColorGray  = lipgloss.Color("241") // Gray
	ColorWhite = lipgloss.Color("255") // White

	// CursorStyle for inputs
	CursorStyle = lipgloss.NewStyle().Foreground(ColorBlue)

	// DocStyle for document layout
	DocStyle = lipgloss.NewStyle().Margin(1, 2)

	// WelcomeStyle for welcome banner
	WelcomeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorBlue)

	// SubtitleStyle
	SubtitleStyle = lipgloss.NewStyle().
			Foreground(ColorGray)

	// KeysStyle for keyboard shortcuts display
	KeysStyle = lipgloss.NewStyle().
			Foreground(ColorGray)

	// ListTitleStyle for list titles
	ListTitleStyle = lipgloss.NewStyle().
			Foreground(ColorBlue).
			Bold(true)

	// ErrorStyle for error messages
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true)

	// SuccessStyle for success messages
	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("46")).
			Bold(true)

	// StatusStyle for status messages
	StatusStyle = lipgloss.NewStyle().
			Foreground(ColorGray)

	// TitleStyle for section titles
	TitleStyle = lipgloss.NewStyle().
			Foreground(ColorBlue).
			Bold(true)

	// HighlightStyle for highlighted text
	HighlightStyle = lipgloss.NewStyle().
			Foreground(ColorBlue).
			Bold(true)

	// SubtleStyle for subtle/muted text
	SubtleStyle = lipgloss.NewStyle().
			Foreground(ColorGray)

	// BoxStyle for bordered content
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBlue).
			Padding(1, 2)

	// InfoStyle for informational messages
	InfoStyle = lipgloss.NewStyle().
			Foreground(ColorGray)

	// MenuItemStyle for menu items
	MenuItemStyle = lipgloss.NewStyle().
			Foreground(ColorWhite)

	// SelectedMenuItemStyle for selected menu items
	SelectedMenuItemStyle = lipgloss.NewStyle().
				Foreground(ColorBlue).
				Bold(true)

	// WarningStyle for warning messages
	WarningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("226")).
			Bold(true)

	// CodeStyle for code snippets
	CodeStyle = lipgloss.NewStyle().
			Foreground(ColorGray)

	// ScrollHintStyle for scroll instructions
	ScrollHintStyle = lipgloss.NewStyle().
			Foreground(ColorGray).
			Italic(true)
)
