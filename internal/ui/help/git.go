package help

// GetGitGuide returns git version control guide
func GetGitGuide() string {
	return headerStyle.Render("🔄 GIT VERSION CONTROL") + "\n\n" +
		textStyle.Render("HOW TO USE:") + "\n" +
		textStyle.Render("  1. Press G from home screen") + "\n" +
		textStyle.Render("  2. Select git operation (1-4)") + "\n" +
		textStyle.Render("  3. Follow prompts") + "\n\n" +
		textStyle.Render("GIT OPERATIONS:") + "\n" +
		textStyle.Render("  1. Initialize Repo - Start version control") + "\n" +
		textStyle.Render("  2. Commit Changes - Save snapshot") + "\n" +
		textStyle.Render("  3. View History - See past commits") + "\n" +
		textStyle.Render("  4. Push to Remote - Sync to GitHub/GitLab") + "\n\n" +
		textStyle.Render("BENEFITS:") + "\n" +
		textStyle.Render("  • Track note changes over time") + "\n" +
		textStyle.Render("  • Revert to previous versions") + "\n" +
		textStyle.Render("  • Collaborate with others") + "\n" +
		textStyle.Render("  • Backup to remote server") + "\n\n" +
		textStyle.Render("WORKFLOW:") + "\n" +
		textStyle.Render("  1. Initialize repo (first time only)") + "\n" +
		textStyle.Render("  2. Make changes to notes") + "\n" +
		textStyle.Render("  3. Commit with descriptive message") + "\n" +
		textStyle.Render("  4. Push to remote (optional)") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
