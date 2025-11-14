package help

// GetExportImportGuide returns export and import guide
func GetExportImportGuide() string {
	return headerStyle.Render("📤📥 EXPORT & IMPORT") + "\n\n" +
		successStyle.Render("EXPORTING:") + "\n" +
		textStyle.Render("  1. Open or select a note") + "\n" +
		textStyle.Render("  2. Press Ctrl+E") + "\n" +
		textStyle.Render("  3. Choose format (1-4)") + "\n" +
		textStyle.Render("  4. File saved to: ~/totion-exports/") + "\n\n" +
		textStyle.Render("EXPORT FORMATS:") + "\n" +
		textStyle.Render("  1. HTML - Web-ready format") + "\n" +
		textStyle.Render("  2. PDF - Print-ready document") + "\n" +
		textStyle.Render("  3. Plain Text - Universal format") + "\n" +
		textStyle.Render("  4. Markdown - Keep original format") + "\n\n" +
		successStyle.Render("IMPORTING:") + "\n" +
		textStyle.Render("  1. Press Ctrl+I from home") + "\n" +
		textStyle.Render("  2. Choose source (1-3)") + "\n" +
		textStyle.Render("  3. Follow on-screen instructions") + "\n\n" +
		textStyle.Render("IMPORT SOURCES:") + "\n" +
		textStyle.Render("  1. Notion - Export from Notion") + "\n" +
		textStyle.Render("  2. Obsidian - Import Obsidian vault") + "\n" +
		textStyle.Render("  3. Plain Text - Copy .md/.txt files") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
