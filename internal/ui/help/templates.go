package help

// GetTemplatesGuide returns templates guide
func GetTemplatesGuide() string {
	return headerStyle.Render("📝 TEMPLATES GUIDE") + "\n\n" +
		textStyle.Render("HOW TO USE:") + "\n" +
		textStyle.Render("  1. Press Ctrl+T from home screen") + "\n" +
		textStyle.Render("  2. Select a template (press 1-7)") + "\n" +
		textStyle.Render("  3. Enter note name") + "\n" +
		textStyle.Render("  4. Template content loads automatically") + "\n" +
		textStyle.Render("  5. Edit and customize as needed") + "\n\n" +
		textStyle.Render("AVAILABLE TEMPLATES:") + "\n" +
		textStyle.Render("  1. Meeting Notes - Structured meeting documentation") + "\n" +
		textStyle.Render("  2. Todo List - Task management with checkboxes") + "\n" +
		textStyle.Render("  3. Journal Entry - Daily reflection format") + "\n" +
		textStyle.Render("  4. Project Plan - Project planning structure") + "\n" +
		textStyle.Render("  5. Code Documentation - Technical doc template") + "\n" +
		textStyle.Render("  6. Research Notes - Academic/research format") + "\n" +
		textStyle.Render("  7. Book Summary - Book review structure") + "\n\n" +
		textStyle.Render("TEMPLATE FEATURES:") + "\n" +
		textStyle.Render("  • Pre-formatted sections") + "\n" +
		textStyle.Render("  • Markdown formatting included") + "\n" +
		textStyle.Render("  • Save time on note structure") + "\n" +
		textStyle.Render("  • Consistent documentation") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
