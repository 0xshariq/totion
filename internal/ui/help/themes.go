package help

// GetThemesGuide returns themes guide
func GetThemesGuide(translate func(string) string) string {
	return headerStyle.Render(translate("🎨 THEMES GUIDE")) + "\n\n" +
		textStyle.Render(translate("HOW TO CHANGE THEMES:")) + "\n" +
		textStyle.Render(translate("  1. Press P from home screen")) + "\n" +
		textStyle.Render(translate("  2. Select a theme (press 1-6)")) + "\n" +
		textStyle.Render(translate("  3. Theme applies immediately")) + "\n" +
		textStyle.Render(translate("  4. Settings persist across sessions")) + "\n\n" +
		textStyle.Render(translate("AVAILABLE THEMES:")) + "\n" +
		textStyle.Render(translate("  1. Ocean - Cool blues and teals")) + "\n" +
		textStyle.Render(translate("  2. Forest - Earthy greens")) + "\n" +
		textStyle.Render(translate("  3. Sunset - Warm oranges and reds")) + "\n" +
		textStyle.Render(translate("  4. Midnight - Dark purple and blue")) + "\n" +
		textStyle.Render(translate("  5. Rose - Soft pinks and purples")) + "\n" +
		textStyle.Render(translate("  6. Monochrome - Classic black and white")) + "\n\n" +
		textStyle.Render(translate("THEME FEATURES:")) + "\n" +
		textStyle.Render(translate("  • Consistent color palette")) + "\n" +
		textStyle.Render(translate("  • Easy on the eyes")) + "\n" +
		textStyle.Render(translate("  • Optimized for terminal")) + "\n" +
		textStyle.Render(translate("  • Professional appearance")) + "\n\n" +
		dimStyle.Render(translate("Press Esc to go back"))
}
