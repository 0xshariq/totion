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
// TranslateFunc is a function that translates text
func GetHelpMenu(translate func(string) string) string {
	return headerStyle.Render(translate("📚 HELP MENU")) + "\n\n" +
		textStyle.Render(translate("Select a topic to learn more:")) + "\n\n" +
		menuStyle.Render("  1. ⌨️  "+translate("Keyboard Shortcuts - Complete key reference")) + "\n" +
		menuStyle.Render("  2. 🚀 "+translate("Getting Started - Quick start guide")) + "\n" +
		menuStyle.Render("  3. 📝 "+translate("Templates Guide - Pre-made note templates")) + "\n" +
		menuStyle.Render("  4. 🎨 "+translate("Themes Guide - Customize appearance")) + "\n" +
		menuStyle.Render("  5. 📤 "+translate("Export & Import - Move notes in/out")) + "\n" +
		menuStyle.Render("  6. 🔄 "+translate("Git Version Control - Track changes")) + "\n" +
		menuStyle.Render("  7. 📊 "+translate("Statistics & Analytics - View insights")) + "\n" +
		menuStyle.Render("  8. ☁️  "+translate("Sync & Backup - Cloud sync setup")) + "\n" +
		menuStyle.Render("  9. 📂 "+translate("Notebooks & Organization - Folder system")) + "\n" +
		menuStyle.Render("  0. 💻 "+translate("Developer Integration - API docs")) + "\n" +
		menuStyle.Render("  T. 🌐 "+translate("UI Translation - Multi-language support")) + "\n\n" +
		dimStyle.Render(translate("Press 1-9, 0, or T to view a topic")+" • "+translate("Press Esc to go back to home"))
}
