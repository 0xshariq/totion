package help

// GetDeveloperGuide returns the developer integration main menu
func GetDeveloperGuide() string {
	return headerStyle.Render(" DEVELOPER INTEGRATION") + "\n\n" +
		textStyle.Render("Choose a topic by pressing the number:") + "\n\n" +
		menuStyle.Render("  1. Core Package Structure") + "\n" +
		menuStyle.Render("  2. Storage & File System API") + "\n" +
		menuStyle.Render("  3. Statistics & Analytics API") + "\n" +
		menuStyle.Render("  4. Export & Import API") + "\n" +
		menuStyle.Render("  5. Templates API") + "\n" +
		menuStyle.Render("  6. Advanced Features (Pinned, Tags, Search)") + "\n" +
		menuStyle.Render("  7. Integration Examples") + "\n" +
		menuStyle.Render("  8. Notebooks API") + "\n" +
		menuStyle.Render("  9. Best Practices") + "\n\n" +
		dimStyle.Render("Press Esc to go back to main help menu")
}

// GetDeveloperCorePackages returns core packages overview
func GetDeveloperCorePackages() string {
	return headerStyle.Render("📦 CORE PACKAGES OVERVIEW") + "\n\n" +
		successStyle.Render("TOTION'S MODULAR ARCHITECTURE:") + "\n" +
		textStyle.Render("Totion is built with a modular, easy-to-integrate architecture.") + "\n" +
		textStyle.Render("Each package handles a specific feature and can be used independently.") + "\n\n" +

		successStyle.Render("MAIN PACKAGES:") + "\n\n" +

		textStyle.Render("📁 internal/storage") + "\n" +
		dimStyle.Render("   Core file management and note CRUD operations") + "\n" +
		dimStyle.Render("   • Create, read, update, delete notes") + "\n" +
		dimStyle.Render("   • Manages vault directory (~/.totion)") + "\n" +
		dimStyle.Render("   • Supports .md and .txt formats") + "\n\n" +

		textStyle.Render("📊 internal/features/stats") + "\n" +
		dimStyle.Render("   Statistics, analytics, and dashboard metrics") + "\n" +
		dimStyle.Render("   • Track note count, word count, character count") + "\n" +
		dimStyle.Render("   • Daily/weekly/monthly trends") + "\n" +
		dimStyle.Render("   • Most active days, writing streaks") + "\n" +
		dimStyle.Render("   • JSON persistence for historical data") + "\n\n" +

		textStyle.Render("📤 internal/features/export") + "\n" +
		dimStyle.Render("   Export notes to multiple formats") + "\n" +
		dimStyle.Render("   • HTML: Beautiful web page format") + "\n" +
		dimStyle.Render("   • PDF: Print-ready documents") + "\n" +
		dimStyle.Render("   • Markdown: Preserve formatting") + "\n" +
		dimStyle.Render("   • Plain Text: Strip all formatting") + "\n\n" +

		textStyle.Render("📋 internal/features/templates") + "\n" +
		dimStyle.Render("   Pre-made note templates system") + "\n" +
		dimStyle.Render("   • Meeting notes, todo lists, journals") + "\n" +
		dimStyle.Render("   • Project plans, code snippets, book notes") + "\n" +
		dimStyle.Render("   • Easy to add custom templates") + "\n\n" +

		textStyle.Render("📌 internal/features/pinned") + "\n" +
		dimStyle.Render("   Pin/favorite notes for quick access") + "\n" +
		dimStyle.Render("   • Pin up to 10 important notes") + "\n" +
		dimStyle.Render("   • Persisted across sessions") + "\n" +
		dimStyle.Render("   • Displayed on home screen") + "\n\n" +

		textStyle.Render("#️⃣  internal/features/tags") + "\n" +
		dimStyle.Render("   Automatic tag extraction and indexing") + "\n" +
		dimStyle.Render("   • Finds #hashtags in notes") + "\n" +
		dimStyle.Render("   • Build tag index for filtering") + "\n" +
		dimStyle.Render("   • Track tag usage frequency") + "\n\n" +

		textStyle.Render("🔍 internal/features/search") + "\n" +
		dimStyle.Render("   Full-text search across all notes") + "\n" +
		dimStyle.Render("   • Fast regex-based search") + "\n" +
		dimStyle.Render("   • Case-insensitive matching") + "\n" +
		dimStyle.Render("   • Search note content and filenames") + "\n\n" +

		textStyle.Render("📅 internal/features/daily") + "\n" +
		dimStyle.Render("   Daily journaling with auto-dated notes") + "\n" +
		dimStyle.Render("   • Creates Daily/YYYY-MM-DD.md files") + "\n" +
		dimStyle.Render("   • Pre-filled template with date/time") + "\n" +
		dimStyle.Render("   • Track daily note count") + "\n\n" +

		textStyle.Render("⚡ internal/features/quick") + "\n" +
		dimStyle.Render("   Quick note scratch pad for temporary notes") + "\n" +
		dimStyle.Render("   • Persistent scratch pad storage") + "\n" +
		dimStyle.Render("   • Convert scratch to permanent note") + "\n" +
		dimStyle.Render("   • Perfect for quick idea capture") + "\n\n" +

		textStyle.Render("📂 internal/notebook") + "\n" +
		dimStyle.Render("   Folder-based notebook organization") + "\n" +
		dimStyle.Render("   • Create, rename, delete notebooks") + "\n" +
		dimStyle.Render("   • Move notes between notebooks") + "\n" +
		dimStyle.Render("   • Get notebook metadata and stats") + "\n\n" +

		successStyle.Render("KEY ARCHITECTURE PRINCIPLES:") + "\n" +
		textStyle.Render("  ✓ Self-contained packages - Each feature is independent") + "\n" +
		textStyle.Render("  ✓ Config directory pattern - All data in ~/.totion/") + "\n" +
		textStyle.Render("  ✓ JSON persistence - Easy integration and data portability") + "\n" +
		textStyle.Render("  ✓ Standard Go patterns - Familiar error handling, interfaces") + "\n" +
		textStyle.Render("  ✓ Thread-safe operations - Safe for concurrent access") + "\n" +
		textStyle.Render("  ✓ Zero external dependencies - Pure Go implementation") + "\n\n" +

		dimStyle.Render("Press Esc to go back to developer menu")
}

// GetDeveloperStorage returns storage API documentation
func GetDeveloperStorage() string {
	return headerStyle.Render("💾 STORAGE & FILE MANAGEMENT API") + "\n\n" +
		successStyle.Render("PACKAGE: internal/storage") + "\n" +
		textStyle.Render("Core API for managing notes and vault directory.") + "\n\n" +

		successStyle.Render("INITIALIZATION:") + "\n" +
		codeStyle.Render("  import \"github.com/0xshariq/totion/internal/storage\"") + "\n" +
		codeStyle.Render("  import \"github.com/0xshariq/totion/internal/models\"") + "\n\n" +
		codeStyle.Render("  storage := storage.New()") + "\n" +
		dimStyle.Render("  // Automatically creates ~/.totion if not exists") + "\n\n" +

		successStyle.Render("CORE METHODS:") + "\n\n" +

		textStyle.Render("CREATE NOTES:") + "\n" +
		codeStyle.Render("  CreateNote(name string, format models.FileFormat) (*os.File, error)") + "\n" +
		dimStyle.Render("  • Creates a new note file in vault") + "\n" +
		dimStyle.Render("  • name: Filename without extension (e.g., \"todo\")") + "\n" +
		dimStyle.Render("  • format: models.FormatMarkdown or models.FormatText") + "\n" +
		dimStyle.Render("  • Returns: Open file handle for writing") + "\n" +
		dimStyle.Render("  • Error: File already exists, permission denied, etc.") + "\n\n" +

		textStyle.Render("LIST NOTES:") + "\n" +
		codeStyle.Render("  ListNotes() ([]models.Note, error)") + "\n" +
		dimStyle.Render("  • Returns all notes in vault directory") + "\n" +
		dimStyle.Render("  • Each Note contains: Name, Path, Format, ModTime, Size") + "\n" +
		dimStyle.Render("  • Excludes hidden files and system files") + "\n" +
		dimStyle.Render("  • Sorted by modification time (newest first)") + "\n\n" +

		textStyle.Render("READ NOTES:") + "\n" +
		codeStyle.Render("  ReadNote(path string) (string, error)") + "\n" +
		dimStyle.Render("  • Reads entire note content from file") + "\n" +
		dimStyle.Render("  • path: Full file path or relative to vault") + "\n" +
		dimStyle.Render("  • Returns: File content as string") + "\n" +
		dimStyle.Render("  • Handles large files efficiently") + "\n\n" +

		textStyle.Render("SAVE NOTES:") + "\n" +
		codeStyle.Render("  SaveNote(file *os.File, content string) error") + "\n" +
		dimStyle.Render("  • Writes content to file and closes it") + "\n" +
		dimStyle.Render("  • Automatically syncs to disk") + "\n" +
		dimStyle.Render("  • Truncates file before writing (overwrites)") + "\n" +
		dimStyle.Render("  • Closes file handle after write") + "\n\n" +

		textStyle.Render("DELETE NOTES:") + "\n" +
		codeStyle.Render("  DeleteNote(path string) error") + "\n" +
		dimStyle.Render("  • Permanently deletes note file") + "\n" +
		dimStyle.Render("  • No trash/recycle bin - cannot undo!") + "\n" +
		dimStyle.Render("  • Use with caution in production") + "\n\n" +

		textStyle.Render("VAULT PATH:") + "\n" +
		codeStyle.Render("  GetVaultPath() string") + "\n" +
		dimStyle.Render("  • Returns ~/.totion directory path") + "\n" +
		dimStyle.Render("  • Use for custom file operations") + "\n\n" +

		successStyle.Render("COMPLETE EXAMPLE:") + "\n" +
		codeStyle.Render("  package main") + "\n\n" +
		codeStyle.Render("  import (") + "\n" +
		codeStyle.Render("    \"log\"") + "\n" +
		codeStyle.Render("    \"github.com/0xshariq/totion/internal/storage\"") + "\n" +
		codeStyle.Render("    \"github.com/0xshariq/totion/internal/models\"") + "\n" +
		codeStyle.Render("  )") + "\n\n" +
		codeStyle.Render("  func main() {") + "\n" +
		codeStyle.Render("    // Initialize storage") + "\n" +
		codeStyle.Render("    s := storage.New()") + "\n\n" +
		codeStyle.Render("    // Create a new markdown note") + "\n" +
		codeStyle.Render("    file, err := s.CreateNote(\"meeting-notes\", models.FormatMarkdown)") + "\n" +
		codeStyle.Render("    if err != nil {") + "\n" +
		codeStyle.Render("      log.Fatal(err)") + "\n" +
		codeStyle.Render("    }") + "\n\n" +
		codeStyle.Render("    // Write content") + "\n" +
		codeStyle.Render("    content := \"# Team Meeting\\n\\n- Discussed Q4 goals\\n- Reviewed roadmap\"") + "\n" +
		codeStyle.Render("    if err := s.SaveNote(file, content); err != nil {") + "\n" +
		codeStyle.Render("      log.Fatal(err)") + "\n" +
		codeStyle.Render("    }") + "\n\n" +
		codeStyle.Render("    // List all notes") + "\n" +
		codeStyle.Render("    notes, err := s.ListNotes()") + "\n" +
		codeStyle.Render("    if err != nil {") + "\n" +
		codeStyle.Render("      log.Fatal(err)") + "\n" +
		codeStyle.Render("    }") + "\n" +
		codeStyle.Render("    log.Printf(\"Total notes: %d\", len(notes))") + "\n" +
		codeStyle.Render("  }") + "\n\n" +

		dimStyle.Render("Press Esc to go back to developer menu")
}
