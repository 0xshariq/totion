package help

// GetExportImportGuide returns export and import guide
func GetExportImportGuide(translate func(string) string) string {
	return headerStyle.Render(translate("📤📥 EXPORT & IMPORT")) + "\n\n" +
		successStyle.Render(translate("EXPORTING:")) + "\n" +
		textStyle.Render(translate("  1. Open or select a note")) + "\n" +
		textStyle.Render(translate("  2. Press Ctrl+E")) + "\n" +
		textStyle.Render(translate("  3. Choose format (1-4)")) + "\n" +
		textStyle.Render(translate("  4. File saved to: ~/totion-exports/")) + "\n\n" +
		textStyle.Render(translate("EXPORT FORMATS:")) + "\n" +
		textStyle.Render(translate("  1. HTML - Web-ready format")) + "\n" +
		textStyle.Render(translate("  2. PDF - Print-ready document")) + "\n" +
		textStyle.Render(translate("  3. Plain Text - Universal format")) + "\n" +
		textStyle.Render(translate("  4. Markdown - Keep original format")) + "\n\n" +
		successStyle.Render(translate("IMPORTING:")) + "\n" +
		textStyle.Render(translate("  1. Press Ctrl+I from home")) + "\n" +
		textStyle.Render(translate("  2. Choose source (1-3)")) + "\n" +
		textStyle.Render(translate("  3. Follow on-screen instructions")) + "\n\n" +
		textStyle.Render(translate("IMPORT SOURCES:")) + "\n" +
		textStyle.Render(translate("  1. Notion - Export from Notion")) + "\n" +
		textStyle.Render(translate("  2. Obsidian - Import Obsidian vault")) + "\n" +
		textStyle.Render(translate("  3. Plain Text - Copy .md/.txt files")) + "\n\n" +
		dimStyle.Render(translate("Press Esc to go back"))
}
