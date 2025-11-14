package help

// GetNotebooksGuide returns notebooks organization guide
func GetNotebooksGuide(translate func(string) string) string {
	return headerStyle.Render(translate("📂 NOTEBOOKS & ORGANIZATION")) + "\n\n" +
		textStyle.Render(translate("HOW TO USE:")) + "\n" +
		textStyle.Render(translate("  1. Press B from home screen")) + "\n" +
		textStyle.Render(translate("  2. Select notebook operation (1-6)")) + "\n" +
		textStyle.Render(translate("  3. Follow prompts")) + "\n\n" +
		textStyle.Render(translate("NOTEBOOK OPERATIONS:")) + "\n" +
		textStyle.Render(translate("  1. Create Notebook - New folder")) + "\n" +
		textStyle.Render(translate("  2. List Notebooks - View all folders")) + "\n" +
		textStyle.Render(translate("  3. Move Note - Organize notes")) + "\n" +
		textStyle.Render(translate("  4. Rename Notebook - Change folder name")) + "\n" +
		textStyle.Render(translate("  5. Delete Notebook - Remove folder")) + "\n" +
		textStyle.Render(translate("  6. Create Note in Notebook - Direct creation")) + "\n\n" +
		textStyle.Render(translate("BENEFITS:")) + "\n" +
		textStyle.Render(translate("  • Keep related notes together")) + "\n" +
		textStyle.Render(translate("  • Hierarchical organization")) + "\n" +
		textStyle.Render(translate("  • Easy to navigate")) + "\n" +
		textStyle.Render(translate("  • Flexible structure")) + "\n\n" +
		textStyle.Render(translate("ORGANIZATION IDEAS:")) + "\n" +
		textStyle.Render(translate("  • By project: Work, Personal, Study")) + "\n" +
		textStyle.Render(translate("  • By type: Meetings, Todos, Journal")) + "\n" +
		textStyle.Render(translate("  • By date: 2025, 2024, Archive")) + "\n" +
		textStyle.Render(translate("  • By topic: Programming, Design, Ideas")) + "\n\n" +
		dimStyle.Render(translate("Press Esc to go back"))
}
