package help

// GetSyncGuide returns sync and backup guide
func GetSyncGuide(translate func(string) string) string {
	return headerStyle.Render(translate("☁️  SYNC & BACKUP")) + "\n\n" +
		textStyle.Render(translate("HOW TO USE:")) + "\n" +
		textStyle.Render(translate("  1. Press Ctrl+Y from home")) + "\n" +
		textStyle.Render(translate("  2. Select sync option (1-4)")) + "\n" +
		textStyle.Render(translate("  3. Configure settings")) + "\n" +
		textStyle.Render(translate("  4. Sync runs automatically")) + "\n\n" +
		textStyle.Render(translate("SYNC OPTIONS:")) + "\n" +
		textStyle.Render(translate("  1. Enable Auto-Sync - Automatic backup")) + "\n" +
		textStyle.Render(translate("  2. Manual Sync - Sync on demand")) + "\n" +
		textStyle.Render(translate("  3. Configure Destination - Set backup location")) + "\n" +
		textStyle.Render(translate("  4. Sync Status - View last sync time")) + "\n\n" +
		textStyle.Render(translate("BACKUP LOCATIONS:")) + "\n" +
		textStyle.Render(translate("  • Dropbox folder")) + "\n" +
		textStyle.Render(translate("  • Google Drive folder")) + "\n" +
		textStyle.Render(translate("  • Custom directory")) + "\n" +
		textStyle.Render(translate("  • External drive")) + "\n\n" +
		textStyle.Render(translate("FEATURES:")) + "\n" +
		textStyle.Render(translate("  • Automatic background sync")) + "\n" +
		textStyle.Render(translate("  • Conflict resolution")) + "\n" +
		textStyle.Render(translate("  • Incremental backups")) + "\n" +
		textStyle.Render(translate("  • Version history")) + "\n\n" +
		dimStyle.Render(translate("Press Esc to go back"))
}
