package help

// GetTemplatesGuide returns templates guide
func GetTemplatesGuide(translate func(string) string) string {
	return headerStyle.Render(translate("📝 TEMPLATES GUIDE")) + "\n\n" +
		textStyle.Render(translate("HOW TO USE:")) + "\n" +
		textStyle.Render(translate("  1. Press Ctrl+T from home screen")) + "\n" +
		textStyle.Render(translate("  2. Select a template (press 1-7)")) + "\n" +
		textStyle.Render(translate("  3. Enter note name")) + "\n" +
		textStyle.Render(translate("  4. Template content loads automatically")) + "\n" +
		textStyle.Render(translate("  5. Edit and customize as needed")) + "\n\n" +
		textStyle.Render(translate("AVAILABLE TEMPLATES:")) + "\n" +
		textStyle.Render(translate("  1. Meeting Notes - Structured meeting documentation")) + "\n" +
		textStyle.Render(translate("  2. Todo List - Task management with checkboxes")) + "\n" +
		textStyle.Render(translate("  3. Journal Entry - Daily reflection format")) + "\n" +
		textStyle.Render(translate("  4. Project Plan - Project planning structure")) + "\n" +
		textStyle.Render(translate("  5. Code Documentation - Technical doc template")) + "\n" +
		textStyle.Render(translate("  6. Research Notes - Academic/research format")) + "\n" +
		textStyle.Render(translate("  7. Book Summary - Book review structure")) + "\n\n" +
		textStyle.Render(translate("TEMPLATE FEATURES:")) + "\n" +
		textStyle.Render(translate("  • Pre-formatted sections")) + "\n" +
		textStyle.Render(translate("  • Markdown formatting included")) + "\n" +
		textStyle.Render(translate("  • Save time on note structure")) + "\n" +
		textStyle.Render(translate("  • Consistent documentation")) + "\n\n" +
		dimStyle.Render(translate("Press Esc to go back"))
}
