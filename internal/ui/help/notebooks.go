package help

// GetNotebooksGuide returns notebooks organization guide
func GetNotebooksGuide() string {
	return headerStyle.Render("📂 NOTEBOOKS & ORGANIZATION") + "\n\n" +
		textStyle.Render("HOW TO USE:") + "\n" +
		textStyle.Render("  1. Press B from home screen") + "\n" +
		textStyle.Render("  2. Select notebook operation (1-6)") + "\n" +
		textStyle.Render("  3. Follow prompts") + "\n\n" +
		textStyle.Render("NOTEBOOK OPERATIONS:") + "\n" +
		textStyle.Render("  1. Create Notebook - New folder") + "\n" +
		textStyle.Render("  2. List Notebooks - View all folders") + "\n" +
		textStyle.Render("  3. Move Note - Organize notes") + "\n" +
		textStyle.Render("  4. Rename Notebook - Change folder name") + "\n" +
		textStyle.Render("  5. Delete Notebook - Remove folder") + "\n" +
		textStyle.Render("  6. Create Note in Notebook - Direct creation") + "\n\n" +
		textStyle.Render("BENEFITS:") + "\n" +
		textStyle.Render("  • Keep related notes together") + "\n" +
		textStyle.Render("  • Hierarchical organization") + "\n" +
		textStyle.Render("  • Easy to navigate") + "\n" +
		textStyle.Render("  • Flexible structure") + "\n\n" +
		textStyle.Render("ORGANIZATION IDEAS:") + "\n" +
		textStyle.Render("  • By project: Work, Personal, Study") + "\n" +
		textStyle.Render("  • By type: Meetings, Todos, Journal") + "\n" +
		textStyle.Render("  • By date: 2025, 2024, Archive") + "\n" +
		textStyle.Render("  • By topic: Programming, Design, Ideas") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
