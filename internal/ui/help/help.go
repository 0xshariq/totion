package help

import (
	"github.com/charmbracelet/lipgloss"
)

// Shared styles for all help files
var (
	headerStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("51")).Bold(true)
	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	dimStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	menuStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("45"))
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true)
	codeStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("228"))
)

// GetHelpMenu returns the main help menu
func GetHelpMenu() string {
	return headerStyle.Render("📚 HELP MENU") + "\n\n" +
		textStyle.Render("Select a topic to learn more:") + "\n\n" +
		menuStyle.Render("  1. ⌨️  Keyboard Shortcuts - Complete key reference") + "\n" +
		menuStyle.Render("  2. 🚀 Getting Started - Quick start guide") + "\n" +
		menuStyle.Render("  3. 📝 Templates Guide - Pre-made note templates") + "\n" +
		menuStyle.Render("  4. 🎨 Themes Guide - Customize appearance") + "\n" +
		menuStyle.Render("  5. 📤 Export & Import - Move notes in/out") + "\n" +
		menuStyle.Render("  6. 🔄 Git Version Control - Track changes") + "\n" +
		menuStyle.Render("  7. 📊 Statistics & Analytics - View insights") + "\n" +
		menuStyle.Render("  8. ☁️  Sync & Backup - Cloud sync setup") + "\n" +
		menuStyle.Render("  9. 📂 Notebooks & Organization - Folder system") + "\n" +
		menuStyle.Render("  0. 💻 Developer Integration - API docs") + "\n" +
		menuStyle.Render("  T. 🌐 UI Translation - Multi-language support") + "\n\n" +
		dimStyle.Render("Press 1-9, 0, or T to view a topic • Press Esc to go back to home")
}
