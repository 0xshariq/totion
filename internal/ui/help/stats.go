package help

// GetStatsGuide returns statistics and analytics guide
func GetStatsGuide() string {
	return headerStyle.Render("📊 STATISTICS & ANALYTICS") + "\n\n" +
		textStyle.Render("HOW TO VIEW:") + "\n" +
		textStyle.Render("  • Press S from home screen") + "\n" +
		textStyle.Render("  • Dashboard shows comprehensive statistics") + "\n" +
		textStyle.Render("  • Updates automatically as you write") + "\n\n" +
		textStyle.Render("METRICS TRACKED:") + "\n" +
		textStyle.Render("  • Total notes created") + "\n" +
		textStyle.Render("  • Total words written") + "\n" +
		textStyle.Render("  • Total characters") + "\n" +
		textStyle.Render("  • Total sentences") + "\n" +
		textStyle.Render("  • Average words per note") + "\n" +
		textStyle.Render("  • Reading time estimate") + "\n" +
		textStyle.Render("  • Writing streak (consecutive days)") + "\n" +
		textStyle.Render("  • Longest streak achieved") + "\n\n" +
		textStyle.Render("DASHBOARD FEATURES:") + "\n" +
		textStyle.Render("  • Weekly activity chart (last 7 days)") + "\n" +
		textStyle.Render("  • Most productive day of week") + "\n" +
		textStyle.Render("  • Top notebooks by note count") + "\n" +
		textStyle.Render("  • ASCII bar charts for visualization") + "\n" +
		textStyle.Render("  • Persistent history tracking") + "\n\n" +
		textStyle.Render("PRODUCTIVITY INSIGHTS:") + "\n" +
		textStyle.Render("  • Track your writing habits") + "\n" +
		textStyle.Render("  • Identify peak productivity times") + "\n" +
		textStyle.Render("  • Set personal goals") + "\n" +
		textStyle.Render("  • Monitor progress over time") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
