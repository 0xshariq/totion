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
menuStyle.Render("  1. Keyboard Shortcuts") + "\n" +
menuStyle.Render("  2. Getting Started") + "\n" +
menuStyle.Render("  3. Templates Guide") + "\n" +
menuStyle.Render("  4. Themes Guide") + "\n" +
menuStyle.Render("  5. Export & Import") + "\n" +
menuStyle.Render("  6. Git Version Control") + "\n" +
menuStyle.Render("  7. Statistics & Analytics") + "\n" +
menuStyle.Render("  8. Sync & Backup") + "\n" +
menuStyle.Render("  9. Notebooks & Organization") + "\n" +
menuStyle.Render("  0. Developer Integration") + "\n\n" +
dimStyle.Render("Press 1-9 or 0 to view a topic") + "\n" +
dimStyle.Render("Press Esc to go back to home")
}
