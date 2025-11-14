package help

// GetGettingStarted returns getting started guide
func GetGettingStarted() string {
	return headerStyle.Render("🚀 GETTING STARTED") + "\n\n" +
		textStyle.Render("1. CREATE YOUR FIRST NOTE:") + "\n" +
		textStyle.Render("   • Press Ctrl+N") + "\n" +
		textStyle.Render("   • Type a name (e.g., 'My First Note')") + "\n" +
		textStyle.Render("   • Press Enter") + "\n" +
		textStyle.Render("   • Choose format: Tab to switch, Enter to create") + "\n" +
		textStyle.Render("   • Start typing your content") + "\n" +
		textStyle.Render("   • Press Ctrl+S to save and close") + "\n\n" +
		textStyle.Render("2. VIEW YOUR NOTES:") + "\n" +
		textStyle.Render("   • Press Ctrl+L to see all notes") + "\n" +
		textStyle.Render("   • Use arrow keys to navigate") + "\n" +
		textStyle.Render("   • Press Enter to open a note") + "\n" +
		textStyle.Render("   • Press / to search") + "\n\n" +
		textStyle.Render("3. DAILY NOTES:") + "\n" +
		textStyle.Render("   • Press Ctrl+D for today's journal") + "\n" +
		textStyle.Render("   • Auto-dated with template") + "\n" +
		textStyle.Render("   • Perfect for daily journaling") + "\n\n" +
		textStyle.Render("4. QUICK NOTES:") + "\n" +
		textStyle.Render("   • Press Ctrl+Q for scratch pad") + "\n" +
		textStyle.Render("   • Jot down quick ideas") + "\n" +
		textStyle.Render("   • Promote to permanent note when ready") + "\n\n" +
		textStyle.Render("5. USE TEMPLATES:") + "\n" +
		textStyle.Render("   • Press Ctrl+T from home") + "\n" +
		textStyle.Render("   • Select a template (1-7)") + "\n" +
		textStyle.Render("   • Template content loads automatically") + "\n\n" +
		textStyle.Render("6. ORGANIZE WITH FOLDERS:") + "\n" +
		textStyle.Render("   • Press B for notebooks menu") + "\n" +
		textStyle.Render("   • Create folders to organize notes") + "\n" +
		textStyle.Render("   • Use # for cross-cutting organization") + "\n\n" +
		successStyle.Render("PRODUCTIVITY TIPS:") + "\n" +
		textStyle.Render("  • Pin favorites with Alt+P (max 10 pinned notes)") + "\n" +
		textStyle.Render("  • Use Alt+F for focus mode (distraction-free writing)") + "\n" +
		textStyle.Render("  • Search across notes with Ctrl+/ (full-text search)") + "\n" +
		textStyle.Render("  • Add # in notes for cross-cutting organization") + "\n" +
		textStyle.Render("  • Use Alt+T to change UI language (works from any screen)") + "\n" +
		textStyle.Render("  • Auto-save runs every 30 seconds in editor") + "\n\n" +
		textStyle.Render("STORAGE:") + "\n" +
		textStyle.Render("  All notes are saved in: ~/.totion/") + "\n" +
		textStyle.Render("  Format: .md (Markdown) or .txt (Plain Text)") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
