package help

// GetSyncGuide returns sync and backup guide
func GetSyncGuide() string {
	return headerStyle.Render("☁️  SYNC & BACKUP") + "\n\n" +
		textStyle.Render("HOW TO USE:") + "\n" +
		textStyle.Render("  1. Press Ctrl+Y from home") + "\n" +
		textStyle.Render("  2. Select sync option (1-4)") + "\n" +
		textStyle.Render("  3. Configure settings") + "\n" +
		textStyle.Render("  4. Sync runs automatically") + "\n\n" +
		textStyle.Render("SYNC OPTIONS:") + "\n" +
		textStyle.Render("  1. Enable Auto-Sync - Automatic backup") + "\n" +
		textStyle.Render("  2. Manual Sync - Sync on demand") + "\n" +
		textStyle.Render("  3. Configure Destination - Set backup location") + "\n" +
		textStyle.Render("  4. Sync Status - View last sync time") + "\n\n" +
		textStyle.Render("BACKUP LOCATIONS:") + "\n" +
		textStyle.Render("  • Dropbox folder") + "\n" +
		textStyle.Render("  • Google Drive folder") + "\n" +
		textStyle.Render("  • Custom directory") + "\n" +
		textStyle.Render("  • External drive") + "\n\n" +
		textStyle.Render("FEATURES:") + "\n" +
		textStyle.Render("  • Automatic background sync") + "\n" +
		textStyle.Render("  • Conflict resolution") + "\n" +
		textStyle.Render("  • Incremental backups") + "\n" +
		textStyle.Render("  • Version history") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
