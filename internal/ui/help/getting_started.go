package help

// GetGettingStarted returns getting started guide
func GetGettingStarted(translate func(string) string) string {
	return headerStyle.Render(translate("🚀 GETTING STARTED")) + "\n\n" +
		textStyle.Render(translate("1. CREATE YOUR FIRST NOTE:")) + "\n" +
		textStyle.Render("   • "+translate("Press Ctrl+N")) + "\n" +
		textStyle.Render("   • "+translate("Type a name (e.g., 'My First Note')")) + "\n" +
		textStyle.Render("   • "+translate("Press Enter")) + "\n" +
		textStyle.Render("   • "+translate("Choose format: Tab to switch, Enter to create")) + "\n" +
		textStyle.Render("   • "+translate("Start typing your content")) + "\n" +
		textStyle.Render("   • "+translate("Press Ctrl+S to save and close")) + "\n\n" +
		textStyle.Render(translate("2. VIEW YOUR NOTES:")) + "\n" +
		textStyle.Render("   • "+translate("Press Ctrl+L to see all notes")) + "\n" +
		textStyle.Render("   • "+translate("Use arrow keys to navigate")) + "\n" +
		textStyle.Render("   • "+translate("Press Enter to open a note")) + "\n" +
		textStyle.Render("   • "+translate("Press / to search")) + "\n\n" +
		textStyle.Render(translate("3. DAILY NOTES:")) + "\n" +
		textStyle.Render("   • "+translate("Press Ctrl+D for today's journal")) + "\n" +
		textStyle.Render("   • "+translate("Auto-dated with template")) + "\n" +
		textStyle.Render("   • "+translate("Perfect for daily journaling")) + "\n\n" +
		textStyle.Render(translate("4. QUICK NOTES:")) + "\n" +
		textStyle.Render("   • "+translate("Press Ctrl+Q for scratch pad")) + "\n" +
		textStyle.Render("   • "+translate("Jot down quick ideas")) + "\n" +
		textStyle.Render("   • "+translate("Promote to permanent note when ready")) + "\n\n" +
		textStyle.Render(translate("5. USE TEMPLATES:")) + "\n" +
		textStyle.Render("   • "+translate("Press Ctrl+T from home")) + "\n" +
		textStyle.Render("   • "+translate("Select a template (1-7)")) + "\n" +
		textStyle.Render("   • "+translate("Template content loads automatically")) + "\n\n" +
		textStyle.Render(translate("6. ORGANIZE WITH FOLDERS:")) + "\n" +
		textStyle.Render("   • "+translate("Press B for notebooks menu")) + "\n" +
		textStyle.Render("   • "+translate("Create folders to organize notes")) + "\n" +
		textStyle.Render("   • "+translate("Use # for cross-cutting organization")) + "\n\n" +
		successStyle.Render(translate("PRODUCTIVITY TIPS:")) + "\n" +
		textStyle.Render("  • "+translate("Pin favorites with Alt+P (max 10 pinned notes)")) + "\n" +
		textStyle.Render("  • "+translate("Use Alt+F for focus mode (distraction-free writing)")) + "\n" +
		textStyle.Render("  • "+translate("Search across notes with Ctrl+/ (full-text search)")) + "\n" +
		textStyle.Render("  • "+translate("Add # in notes for cross-cutting organization")) + "\n" +
		textStyle.Render("  • "+translate("Use Alt+T to change UI language (works from any screen)")) + "\n" +
		textStyle.Render("  • "+translate("Auto-save runs every 30 seconds in editor")) + "\n\n" +
		textStyle.Render(translate("STORAGE:")) + "\n" +
		textStyle.Render("  "+translate("All notes are saved in: ~/.totion/")) + "\n" +
		textStyle.Render("  "+translate("Format: .md (Markdown) or .txt (Plain Text)")) + "\n\n" +
		dimStyle.Render(translate("Press Esc to go back"))
}
