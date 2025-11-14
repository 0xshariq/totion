package help

// GetTranslationGuide returns translation feature guide
func GetTranslationGuide(translate func(string) string) string {
	return headerStyle.Render(translate("🌐 UI TRANSLATION GUIDE")) + "\n\n" +
		successStyle.Render("WHAT GETS TRANSLATED:") + "\n" +
		textStyle.Render("  ✓ All UI elements (menus, buttons, labels)") + "\n" +
		textStyle.Render("  ✓ Help sections and documentation") + "\n" +
		textStyle.Render("  ✓ Status messages and notifications") + "\n" +
		textStyle.Render("  ✓ Keyboard shortcut descriptions") + "\n" +
		textStyle.Render("  ✗ Your note content (stays unchanged)") + "\n" +
		textStyle.Render("  ✗ File names and paths") + "\n\n" +
		headerStyle.Render("HOW TO USE:") + "\n" +
		textStyle.Render("  1. Press Alt+T from ANY screen") + "\n" +
		textStyle.Render("     (homepage, editor, list view, help, etc.)") + "\n" +
		textStyle.Render("  2. Use ↑↓ arrow keys to select language") + "\n" +
		textStyle.Render("  3. Press Enter to apply") + "\n" +
		textStyle.Render("  4. Entire UI translates instantly") + "\n" +
		textStyle.Render("  5. Press Esc to cancel") + "\n\n" +
		headerStyle.Render("SUPPORTED LANGUAGES:") + "\n" +
		menuStyle.Render("  🇪🇸 Spanish (Español)") + "\n" +
		menuStyle.Render("  🇫🇷 French (Français)") + "\n" +
		menuStyle.Render("  🇩🇪 German (Deutsch)") + "\n" +
		menuStyle.Render("  🇯🇵 Japanese (日本語)") + "\n" +
		menuStyle.Render("  🇨🇳 Chinese (中文)") + "\n" +
		menuStyle.Render("  🇰🇷 Korean (한국어)") + "\n" +
		menuStyle.Render("  🇵🇹 Portuguese (Português)") + "\n" +
		menuStyle.Render("  🇮🇹 Italian (Italiano)") + "\n" +
		menuStyle.Render("  🇷🇺 Russian (Русский)") + "\n\n" +
		headerStyle.Render("SETUP INSTRUCTIONS:") + "\n" +
		textStyle.Render("  1. Get a free API key from: https://lingo.dev") + "\n" +
		textStyle.Render("  2. Create .env file in project root") + "\n" +
		textStyle.Render("  3. Add: LINGODOTDEV_API_KEY=your_api_key_here") + "\n" +
		textStyle.Render("  4. (Optional) Set default: LINGO_DEFAULT_LOCALE=es") + "\n" +
		textStyle.Render("  5. Restart Totion") + "\n\n" +
		codeStyle.Render("  Example .env file:") + "\n" +
		codeStyle.Render("  LINGODOTDEV_API_KEY=sk_1234567890abcdef") + "\n" +
		codeStyle.Render("  LINGO_DEFAULT_LOCALE=fr") + "\n\n" +
		headerStyle.Render("TROUBLESHOOTING:") + "\n" +
		textStyle.Render("  • Alt+T does nothing?") + "\n" +
		textStyle.Render("    → Check .env file exists with valid API key") + "\n" +
		textStyle.Render("    → Verify .env is in project root directory") + "\n" +
		textStyle.Render("    → Restart application after adding .env") + "\n\n" +
		textStyle.Render("  • Translation fails?") + "\n" +
		textStyle.Render("    → Check internet connection (API requires online access)") + "\n" +
		textStyle.Render("    → Verify API key is valid (check lingo.dev dashboard)") + "\n" +
		textStyle.Render("    → Check API quota/limits") + "\n\n" +
		textStyle.Render("  • Wrong language selected?") + "\n" +
		textStyle.Render("    → Press Alt+T again and choose different language") + "\n" +
		textStyle.Render("    → Set LINGO_DEFAULT_LOCALE in .env for default") + "\n\n" +
		headerStyle.Render("FEATURES:") + "\n" +
		successStyle.Render("  ✨ Instant translation (< 1 second)") + "\n" +
		successStyle.Render("  ✨ Works from any screen") + "\n" +
		successStyle.Render("  ✨ Preserves formatting and emojis") + "\n" +
		successStyle.Render("  ✨ High-quality AI translation") + "\n" +
		successStyle.Render("  ✨ Remembers your choice per session") + "\n\n" +
		headerStyle.Render("PRIVACY & NOTES:") + "\n" +
		textStyle.Render("  • Only UI text is sent to Lingo.dev API") + "\n" +
		textStyle.Render("  • Your note content is NEVER translated or sent") + "\n" +
		textStyle.Render("  • Notes remain in the language you write them") + "\n" +
		textStyle.Render("  • Translation is for UI accessibility only") + "\n" +
		textStyle.Render("  • Requires internet connection") + "\n" +
		textStyle.Render("  • Free tier available at lingo.dev") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
