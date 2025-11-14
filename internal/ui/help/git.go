package help

// GetGitGuide returns git version control guide
func GetGitGuide(translate func(string) string) string {
	return headerStyle.Render(translate("🔄 GIT VERSION CONTROL")) + "\n\n" +
		textStyle.Render(translate("HOW TO USE:")) + "\n" +
		textStyle.Render(translate("  1. Press G from home screen")) + "\n" +
		textStyle.Render(translate("  2. Select git operation (1-4)")) + "\n" +
		textStyle.Render(translate("  3. Follow prompts")) + "\n\n" +
		textStyle.Render(translate("GIT OPERATIONS:")) + "\n" +
		textStyle.Render(translate("  1. Initialize Repo - Start version control")) + "\n" +
		textStyle.Render(translate("  2. Commit Changes - Save snapshot")) + "\n" +
		textStyle.Render(translate("  3. View History - See past commits")) + "\n" +
		textStyle.Render(translate("  4. Push to Remote - Sync to GitHub/GitLab")) + "\n\n" +
		textStyle.Render(translate("BENEFITS:")) + "\n" +
		textStyle.Render(translate("  • Track note changes over time")) + "\n" +
		textStyle.Render(translate("  • Revert to previous versions")) + "\n" +
		textStyle.Render(translate("  • Collaborate with others")) + "\n" +
		textStyle.Render(translate("  • Backup to remote server")) + "\n\n" +
		textStyle.Render(translate("WORKFLOW:")) + "\n" +
		textStyle.Render(translate("  1. Initialize repo (first time only)")) + "\n" +
		textStyle.Render(translate("  2. Make changes to notes")) + "\n" +
		textStyle.Render(translate("  3. Commit with descriptive message")) + "\n" +
		textStyle.Render(translate("  4. Push to remote (optional)")) + "\n\n" +
		dimStyle.Render(translate("Press Esc to go back"))
}
