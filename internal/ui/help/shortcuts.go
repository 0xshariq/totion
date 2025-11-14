package help

// GetKeyboardShortcuts returns keyboard shortcuts help
func GetKeyboardShortcuts() string {
	return headerStyle.Render("⌨️  KEYBOARD SHORTCUTS") + "\n\n" +
		textStyle.Render("CREATING & EDITING:") + "\n" +
		textStyle.Render("  Ctrl+N      Create new note") + "\n" +
		textStyle.Render("  Ctrl+L      List all notes") + "\n" +
		textStyle.Render("  Ctrl+D      Daily note (home) / Delete (list)") + "\n" +
		textStyle.Render("  Ctrl+Q      Quick note (scratch pad)") + "\n" +
		textStyle.Render("  Ctrl+S      Save and close editor") + "\n" +
		textStyle.Render("  Enter       Open/Create note") + "\n" +
		textStyle.Render("  Esc         Go back / Cancel") + "\n\n" +
		textStyle.Render("EDITOR FEATURES:") + "\n" +
		textStyle.Render("  Alt+F       Toggle focus mode (distraction-free)") + "\n" +
		textStyle.Render("  Alt+T       Translate note to another language") + "\n" +
		textStyle.Render("  Alt+P       Pin/unpin current note (max 10)") + "\n" +
		textStyle.Render("  Alt+L       Wiki linking help") + "\n" +
		textStyle.Render("  Ctrl+S      Save & close (auto-save enabled)") + "\n\n" +
		textStyle.Render("SEARCH & ORGANIZATION:") + "\n" +
		textStyle.Render("  Ctrl+/      Full-text search across notes") + "\n" +
		textStyle.Render("  #tag        Type tags in notes (e.g., #work)") + "\n" +
		textStyle.Render("  /           Search/filter notes in list") + "\n" +
		textStyle.Render("  B           Notebooks (folder organization)") + "\n\n" +
		textStyle.Render("NAVIGATION:") + "\n" +
		textStyle.Render("  ↑ ↓         Navigate lists") + "\n" +
		textStyle.Render("  Tab         Switch format (Markdown/Text)") + "\n\n" +
		textStyle.Render("FEATURES:") + "\n" +
		textStyle.Render("  Ctrl+T      Templates menu") + "\n" +
		textStyle.Render("  P           Themes menu") + "\n" +
		textStyle.Render("  Alt+E       Export note") + "\n" +
		textStyle.Render("  Alt+I       Import notes") + "\n" +
		textStyle.Render("  S           View statistics & dashboard") + "\n" +
		textStyle.Render("  G           Git operations") + "\n" +
		textStyle.Render("  Alt+Y       Sync & backup") + "\n\n" +
		textStyle.Render("GENERAL:") + "\n" +
		textStyle.Render("  Ctrl+H / ?  Show this help") + "\n" +
		textStyle.Render("  Q           Quit application") + "\n\n" +
		successStyle.Render("FOCUS MODE:") + "\n" +
		textStyle.Render("  • Press Alt+F in editor for distraction-free writing") + "\n" +
		textStyle.Render("  • Minimal UI with only editor and word count") + "\n" +
		textStyle.Render("  • Press Alt+F again or Esc to exit") + "\n\n" +
		successStyle.Render("PINNED NOTES:") + "\n" +
		textStyle.Render("  • Press Alt+P in editor to pin/unpin") + "\n" +
		textStyle.Render("  • Pinned notes show 📌 indicator") + "\n" +
		textStyle.Render("  • View pinned notes on home screen") + "\n" +
		textStyle.Render("  • Maximum 10 pinned notes allowed") + "\n\n" +
		successStyle.Render("🌐 TRANSLATION:") + "\n" +
		textStyle.Render("  • Press Alt+T in editor to translate note") + "\n" +
		textStyle.Render("  • Support for 9 languages: ES, FR, DE, JA, ZH, KO, PT, IT, RU") + "\n" +
		textStyle.Render("  • Preserves markdown formatting") + "\n" +
		textStyle.Render("  • Requires LINGODOTDEV_API_KEY (get free key at lingo.dev)") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
