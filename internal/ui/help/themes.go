package help

// GetThemesGuide returns themes guide
func GetThemesGuide() string {
	return headerStyle.Render("🎨 THEMES GUIDE") + "\n\n" +
		textStyle.Render("HOW TO CHANGE THEMES:") + "\n" +
		textStyle.Render("  1. Press P from home screen") + "\n" +
		textStyle.Render("  2. Select a theme (press 1-6)") + "\n" +
		textStyle.Render("  3. Theme applies immediately") + "\n" +
		textStyle.Render("  4. Settings persist across sessions") + "\n\n" +
		textStyle.Render("AVAILABLE THEMES:") + "\n" +
		textStyle.Render("  1. Ocean - Cool blues and teals") + "\n" +
		textStyle.Render("  2. Forest - Earthy greens") + "\n" +
		textStyle.Render("  3. Sunset - Warm oranges and reds") + "\n" +
		textStyle.Render("  4. Midnight - Dark purple and blue") + "\n" +
		textStyle.Render("  5. Rose - Soft pinks and purples") + "\n" +
		textStyle.Render("  6. Monochrome - Classic black and white") + "\n\n" +
		textStyle.Render("THEME FEATURES:") + "\n" +
		textStyle.Render("  • Consistent color palette") + "\n" +
		textStyle.Render("  • Easy on the eyes") + "\n" +
		textStyle.Render("  • Optimized for terminal") + "\n" +
		textStyle.Render("  • Professional appearance") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
