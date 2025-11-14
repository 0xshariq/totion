package help

// GetStatsGuide returns statistics and analytics guide
func GetStatsGuide(translate func(string) string) string {
	return headerStyle.Render(translate("📊 STATISTICS & ANALYTICS")) + "\n\n" +
		textStyle.Render(translate("HOW TO VIEW:")) + "\n" +
		textStyle.Render(translate("  • Press S from home screen")) + "\n" +
		textStyle.Render(translate("  • Dashboard shows comprehensive statistics")) + "\n" +
		textStyle.Render(translate("  • Updates automatically as you write")) + "\n\n" +
		textStyle.Render(translate("METRICS TRACKED:")) + "\n" +
		textStyle.Render(translate("  • Total notes created")) + "\n" +
		textStyle.Render(translate("  • Total words written")) + "\n" +
		textStyle.Render(translate("  • Total characters")) + "\n" +
		textStyle.Render(translate("  • Total sentences")) + "\n" +
		textStyle.Render(translate("  • Average words per note")) + "\n" +
		textStyle.Render(translate("  • Reading time estimate")) + "\n" +
		textStyle.Render(translate("  • Writing streak (consecutive days)")) + "\n" +
		textStyle.Render(translate("  • Longest streak achieved")) + "\n\n" +
		textStyle.Render(translate("DASHBOARD FEATURES:")) + "\n" +
		textStyle.Render(translate("  • Weekly activity chart (last 7 days)")) + "\n" +
		textStyle.Render(translate("  • Most productive day of week")) + "\n" +
		textStyle.Render(translate("  • Top notebooks by note count")) + "\n" +
		textStyle.Render(translate("  • ASCII bar charts for visualization")) + "\n" +
		textStyle.Render(translate("  • Persistent history tracking")) + "\n\n" +
		textStyle.Render(translate("PRODUCTIVITY INSIGHTS:")) + "\n" +
		textStyle.Render(translate("  • Track your writing habits")) + "\n" +
		textStyle.Render(translate("  • Identify peak productivity times")) + "\n" +
		textStyle.Render(translate("  • Set personal goals")) + "\n" +
		textStyle.Render(translate("  • Monitor progress over time")) + "\n\n" +
		dimStyle.Render(translate("Press Esc to go back"))
}
