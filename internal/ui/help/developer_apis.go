package help

// GetDeveloperStats returns statistics API documentation
func GetDeveloperStats() string {
	return headerStyle.Render("📊 STATISTICS & ANALYTICS API") + "\n\n" +
		successStyle.Render("PACKAGE: internal/features/stats") + "\n\n" +
		textStyle.Render("INITIALIZATION:") + "\n" +
		codeStyle.Render("  // Basic stats manager") + "\n" +
		codeStyle.Render("  sm := stats.NewStatsManager()") + "\n\n" +
		codeStyle.Render("  // With persistence") + "\n" +
		codeStyle.Render("  sm := stats.NewStatsManagerWithConfig(configDir)") + "\n\n" +
		textStyle.Render("CORE METHODS:") + "\n" +
		dimStyle.Render("  • Calculate(content string) Statistics") + "\n" +
		dimStyle.Render("    Get word count, reading time, characters, sentences") + "\n\n" +
		dimStyle.Render("  • RecordStats(date time.Time, words, notes int)") + "\n" +
		dimStyle.Render("    Save writing activity for dashboard") + "\n\n" +
		dimStyle.Render("  • GetLongestStreak() int") + "\n" +
		dimStyle.Render("    Get longest consecutive writing days") + "\n\n" +
		dimStyle.Render("  • GetWeeklyActivity() map[string]int") + "\n" +
		dimStyle.Render("    Last 7 days activity breakdown") + "\n\n" +
		dimStyle.Render("  • BuildDashboard(...) DashboardData") + "\n" +
		dimStyle.Render("    Generate comprehensive analytics") + "\n\n" +
		dimStyle.Render("  • RenderDashboard(data DashboardData) string") + "\n" +
		dimStyle.Render("    Create ASCII visualization") + "\n\n" +
		successStyle.Render("EXAMPLE:") + "\n" +
		codeStyle.Render("  sm := stats.NewStatsManager()") + "\n" +
		codeStyle.Render("  stats := sm.Calculate(noteContent)") + "\n" +
		codeStyle.Render("  fmt.Printf(\"Words: %d\\n\", stats.WordCount)") + "\n" +
		codeStyle.Render("  fmt.Printf(\"Reading time: %s\\n\", stats.ReadingTime)") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}

// GetDeveloperExport returns export API documentation
func GetDeveloperExport() string {
	return headerStyle.Render("📤 EXPORT & IMPORT API") + "\n\n" +
		successStyle.Render("PACKAGE: internal/features/export") + "\n\n" +
		textStyle.Render("EXPORT API:") + "\n" +
		codeStyle.Render("  exp := export.NewExporter()") + "\n\n" +
		dimStyle.Render("  • ExportToHTML(note, outputPath) error") + "\n" +
		dimStyle.Render("    Convert note to styled HTML") + "\n\n" +
		dimStyle.Render("  • ExportToPDF(note, outputPath) error") + "\n" +
		dimStyle.Render("    Generate PDF document") + "\n\n" +
		dimStyle.Render("  • ExportToMarkdown(note, outputPath) error") + "\n" +
		dimStyle.Render("    Export as standalone MD file") + "\n\n" +
		dimStyle.Render("  • ExportToText(note, outputPath) error") + "\n" +
		dimStyle.Render("    Convert to plain text") + "\n\n" +
		successStyle.Render("IMPORT API:") + "\n" +
		successStyle.Render("PACKAGE: internal/features/import") + "\n\n" +
		dimStyle.Render("  • ImportFromNotion(exportPath) ([]Note, error)") + "\n" +
		dimStyle.Render("    Import Notion export JSON") + "\n\n" +
		dimStyle.Render("  • ImportFromObsidian(vaultPath) ([]Note, error)") + "\n" +
		dimStyle.Render("    Import Obsidian vault") + "\n\n" +
		successStyle.Render("EXAMPLE:") + "\n" +
		codeStyle.Render("  exp := export.NewExporter()") + "\n" +
		codeStyle.Render("  err := exp.ExportToHTML(note, \"output.html\")") + "\n" +
		codeStyle.Render("  if err != nil { log.Fatal(err) }") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}

// GetDeveloperTemplates returns templates API documentation
func GetDeveloperTemplates() string {
	return headerStyle.Render("📝 TEMPLATES & THEMES API") + "\n\n" +
		successStyle.Render("PACKAGE: internal/features/templates") + "\n\n" +
		textStyle.Render("TEMPLATE MANAGEMENT:") + "\n" +
		codeStyle.Render("  tm := templates.NewTemplateManager()") + "\n\n" +
		dimStyle.Render("  • GetTemplate(name string) (string, error)") + "\n" +
		dimStyle.Render("    Retrieve template content") + "\n\n" +
		dimStyle.Render("  • SaveCustomTemplate(name, content) error") + "\n" +
		dimStyle.Render("    Create custom template") + "\n\n" +
		dimStyle.Render("  • ListTemplates() []string") + "\n" +
		dimStyle.Render("    Get all available templates") + "\n\n" +
		textStyle.Render("BUILT-IN TEMPLATES:") + "\n" +
		dimStyle.Render("  • meeting - Meeting notes structure") + "\n" +
		dimStyle.Render("  • todo - Task list with checkboxes") + "\n" +
		dimStyle.Render("  • journal - Daily journal format") + "\n" +
		dimStyle.Render("  • project - Project plan template") + "\n" +
		dimStyle.Render("  • code - Code documentation") + "\n" +
		dimStyle.Render("  • research - Research notes") + "\n" +
		dimStyle.Render("  • book - Book summary") + "\n\n" +
		successStyle.Render("EXAMPLE:") + "\n" +
		codeStyle.Render("  tm := templates.NewTemplateManager()") + "\n" +
		codeStyle.Render("  content, err := tm.GetTemplate(\"meeting\")") + "\n" +
		codeStyle.Render("  if err != nil { log.Fatal(err) }") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
