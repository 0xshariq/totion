package help

// GetDeveloperDaily returns daily & quick notes API
func GetDeveloperDaily() string {
	return headerStyle.Render("📅 DAILY & QUICK NOTES API") + "\n\n" +
		successStyle.Render("DAILY NOTES:") + "\n" +
		successStyle.Render("PACKAGE: internal/features/daily") + "\n\n" +
		codeStyle.Render("  dm := daily.NewDailyManager(vaultDir)") + "\n\n" +
		dimStyle.Render("  • GetTodayNotePath() string") + "\n" +
		dimStyle.Render("    Returns: ~/.totion/Daily/YYYY-MM-DD.md") + "\n\n" +
		dimStyle.Render("  • CreateTodayNote() (*os.File, error)") + "\n" +
		dimStyle.Render("    Creates or opens today's note with template") + "\n\n" +
		dimStyle.Render("  • TodayNoteExists() bool") + "\n" +
		dimStyle.Render("    Check if today's note exists") + "\n\n" +
		dimStyle.Render("  • GetDailyNotesCount() int") + "\n" +
		dimStyle.Render("    Total daily notes created") + "\n\n" +
		successStyle.Render("QUICK NOTES:") + "\n" +
		successStyle.Render("PACKAGE: internal/features/quick") + "\n\n" +
		codeStyle.Render("  qm := quick.NewQuickNoteManager(vaultDir)") + "\n\n" +
		dimStyle.Render("  • LoadScratch() (string, error)") + "\n" +
		dimStyle.Render("    Load scratch pad content") + "\n\n" +
		dimStyle.Render("  • SaveScratch(content) error") + "\n" +
		dimStyle.Render("    Save scratch pad content") + "\n\n" +
		dimStyle.Render("  • PromoteToNote(vaultDir, name) error") + "\n" +
		dimStyle.Render("    Convert scratch to permanent note") + "\n\n" +
		dimStyle.Render("  • ClearScratch() error") + "\n" +
		dimStyle.Render("    Reset scratch pad") + "\n\n" +
		successStyle.Render("EXAMPLE:") + "\n" +
		codeStyle.Render("  // Daily note") + "\n" +
		codeStyle.Render("  dm := daily.NewDailyManager(vault)") + "\n" +
		codeStyle.Render("  file, _ := dm.CreateTodayNote()") + "\n" +
		codeStyle.Render("  // Note has pre-filled template") + "\n\n" +
		codeStyle.Render("  // Quick note") + "\n" +
		codeStyle.Render("  qm := quick.NewQuickNoteManager(vault)") + "\n" +
		codeStyle.Render("  content, _ := qm.LoadScratch()") + "\n" +
		codeStyle.Render("  // Edit and save...") + "\n" +
		codeStyle.Render("  qm.PromoteToNote(vault, \"My Idea\")") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}

// GetDeveloperNotebooks returns notebooks API
func GetDeveloperNotebooks() string {
	return headerStyle.Render("📂 NOTEBOOKS ORGANIZATION API") + "\n\n" +
		successStyle.Render("PACKAGE: internal/notebook") + "\n\n" +
		textStyle.Render("INITIALIZATION:") + "\n" +
		codeStyle.Render("  nm := notebook.NewNotebookManager(vaultDir)") + "\n\n" +
		textStyle.Render("CORE METHODS:") + "\n" +
		dimStyle.Render("  • CreateNotebook(name string) error") + "\n" +
		dimStyle.Render("  • ListNotebooks() ([]string, error)") + "\n" +
		dimStyle.Render("  • GetNotebookInfo(path) (NotebookInfo, error)") + "\n" +
		dimStyle.Render("  • RenameNotebook(old, new string) error") + "\n" +
		dimStyle.Render("  • DeleteNotebook(name string) error") + "\n" +
		dimStyle.Render("  • MoveNoteToNotebook(notePath, notebook) error") + "\n\n" +
		successStyle.Render("EXAMPLE:") + "\n" +
		codeStyle.Render("  nm := notebook.NewNotebookManager(vault)") + "\n" +
		codeStyle.Render("  nm.CreateNotebook(\"Work\")") + "\n" +
		codeStyle.Render("  nm.MoveNoteToNotebook(\"todo.md\", \"Work\")") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}

// GetDeveloperBestPractices returns integration best practices
func GetDeveloperBestPractices() string {
	return headerStyle.Render("💡 INTEGRATION BEST PRACTICES") + "\n\n" +
		successStyle.Render("1. INITIALIZATION:") + "\n" +
		textStyle.Render("  • Use ~/.totion as config directory") + "\n" +
		textStyle.Render("  • Create directory if not exists") + "\n\n" +
		successStyle.Render("2. ERROR HANDLING:") + "\n" +
		textStyle.Render("  • All methods return Go-style errors") + "\n" +
		textStyle.Render("  • Always check err != nil") + "\n\n" +
		successStyle.Render("3. PERSISTENCE:") + "\n" +
		textStyle.Render("  • .stats.json - Statistics history") + "\n" +
		textStyle.Render("  • .pinned_notes.json - Pinned notes") + "\n" +
		textStyle.Render("  • .tags.json - Tag index") + "\n\n" +
		successStyle.Render("4. THREAD SAFETY:") + "\n" +
		textStyle.Render("  • All managers are thread-safe") + "\n" +
		textStyle.Render("  • JSON operations are atomic") + "\n\n" +
		successStyle.Render("5. IMPORT PATHS:") + "\n" +
		codeStyle.Render("  github.com/0xshariq/totion/internal/storage") + "\n" +
		codeStyle.Render("  github.com/0xshariq/totion/internal/features/*") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
