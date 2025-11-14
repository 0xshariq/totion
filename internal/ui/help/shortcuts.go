package help

// GetKeyboardShortcuts returns keyboard shortcuts help
func GetKeyboardShortcuts(translate func(string) string) string {
	return headerStyle.Render(translate("⌨️  KEYBOARD SHORTCUTS")) + "\n\n" +
		textStyle.Render(translate("CREATING & EDITING:")) + "\n" +
		textStyle.Render(translate("  Ctrl+N      Create new note")) + "\n" +
		textStyle.Render(translate("  Ctrl+L      List all notes")) + "\n" +
		textStyle.Render(translate("  Ctrl+D      Daily note (home) / Delete (list)")) + "\n" +
		textStyle.Render(translate("  Ctrl+Q      Quick note (scratch pad)")) + "\n" +
		textStyle.Render(translate("  Ctrl+S      Save and close editor")) + "\n" +
		textStyle.Render(translate("  Enter       Open/Create note")) + "\n" +
		textStyle.Render(translate("  Esc         Go back / Cancel")) + "\n\n" +
		textStyle.Render(translate("EDITOR FEATURES:")) + "\n" +
		textStyle.Render(translate("  Alt+F       Toggle focus mode (distraction-free)")) + "\n" +
		textStyle.Render(translate("  Alt+T       Translate note to another language")) + "\n" +
		textStyle.Render(translate("  Alt+P       Pin/unpin current note (max 10)")) + "\n" +
		textStyle.Render(translate("  Alt+L       Wiki linking help")) + "\n" +
		textStyle.Render(translate("  Ctrl+S      Save & close (auto-save enabled)")) + "\n\n" +
		textStyle.Render(translate("SEARCH & ORGANIZATION:")) + "\n" +
		textStyle.Render(translate("  Ctrl+/      Full-text search across notes")) + "\n" +
		textStyle.Render(translate("  #           Type tags in notes (e.g., #work)")) + "\n" +
		textStyle.Render(translate("  /           Search/filter notes in list")) + "\n" +
		textStyle.Render(translate("  B           Notebooks (folder organization)")) + "\n\n" +
		textStyle.Render(translate("NAVIGATION:")) + "\n" +
		textStyle.Render(translate("  ↑ ↓         Navigate lists")) + "\n" +
		textStyle.Render(translate("  Tab         Switch format (Markdown/Text)")) + "\n\n" +
		textStyle.Render(translate("FEATURES:")) + "\n" +
		textStyle.Render(translate("  Ctrl+T      Templates menu")) + "\n" +
		textStyle.Render(translate("  P           Themes menu")) + "\n" +
		textStyle.Render(translate("  Alt+E       Export note")) + "\n" +
		textStyle.Render(translate("  Alt+I       Import notes")) + "\n" +
		textStyle.Render(translate("  S           View statistics & dashboard")) + "\n" +
		textStyle.Render(translate("  G           Git operations")) + "\n" +
		textStyle.Render(translate("  Alt+Y       Sync & backup")) + "\n\n" +
		textStyle.Render(translate("GENERAL:")) + "\n" +
		textStyle.Render(translate("  Ctrl+H / ?  Show this help")) + "\n" +
		textStyle.Render(translate("  Q           Quit application")) + "\n\n" +
		successStyle.Render(translate("FOCUS MODE:")) + "\n" +
		textStyle.Render(translate("  • Press Alt+F in editor for distraction-free writing")) + "\n" +
		textStyle.Render(translate("  • Minimal UI with only editor and word count")) + "\n" +
		textStyle.Render(translate("  • Press Alt+F again or Esc to exit")) + "\n\n" +
		successStyle.Render(translate("PINNED NOTES:")) + "\n" +
		textStyle.Render(translate("  • Press Alt+P in editor to pin/unpin")) + "\n" +
		textStyle.Render(translate("  • Pinned notes show 📌 indicator")) + "\n" +
		textStyle.Render(translate("  • View pinned notes on home screen")) + "\n" +
		textStyle.Render(translate("  • Maximum 10 pinned notes allowed")) + "\n\n" +
		successStyle.Render(translate("🌐 UI TRANSLATION:")) + "\n" +
		textStyle.Render(translate("  • Press Alt+T from ANY screen (not just homepage)")) + "\n" +
		textStyle.Render(translate("  • Translates entire UI: menus, labels, help text")) + "\n" +
		textStyle.Render(translate("  • Your notes stay in original language (not translated)")) + "\n" +
		textStyle.Render(translate("  • Languages: Spanish, French, German, Japanese, Chinese,")) + "\n" +
		textStyle.Render(translate("    Korean, Portuguese, Italian, Russian")) + "\n" +
		textStyle.Render(translate("  • Setup: Add LINGODOTDEV_API_KEY to .env file")) + "\n" +
		textStyle.Render(translate("  • Get free API key at: https://lingo.dev")) + "\n" +
		textStyle.Render(translate("  • Set default language: LINGO_DEFAULT_LOCALE=es (optional)")) + "\n\n" +
		dimStyle.Render(translate("Press Esc to go back"))
}
