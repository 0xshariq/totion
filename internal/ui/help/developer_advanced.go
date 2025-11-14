package help

// GetDeveloperAdvanced returns advanced features API
func GetDeveloperAdvanced() string {
	return headerStyle.Render("🚀 ADVANCED FEATURES API") + "\n\n" +
		successStyle.Render("PINNED NOTES:") + "\n" +
		successStyle.Render("PACKAGE: internal/features/pinned") + "\n\n" +
		codeStyle.Render("  pm := pinned.NewPinnedManager(configDir)") + "\n\n" +
		dimStyle.Render("  • Pin(path, name) error - Pin a note") + "\n" +
		dimStyle.Render("  • Unpin(path) error - Unpin note") + "\n" +
		dimStyle.Render("  • Toggle(path, name) (bool, error)") + "\n" +
		dimStyle.Render("  • IsPinned(path) bool") + "\n" +
		dimStyle.Render("  • GetPinned() []PinnedNote") + "\n\n" +
		successStyle.Render("TAG SYSTEM:") + "\n" +
		successStyle.Render("PACKAGE: internal/features/tags") + "\n\n" +
		codeStyle.Render("  tm := tags.NewTagManager(vaultDir, configDir)") + "\n\n" +
		dimStyle.Render("  • ExtractTags(content) []string") + "\n" +
		dimStyle.Render("    Extract #tags from text") + "\n\n" +
		dimStyle.Render("  • RebuildIndex() error") + "\n" +
		dimStyle.Render("    Scan all notes and build index") + "\n\n" +
		dimStyle.Render("  • GetNotesByTag(tag) []string") + "\n" +
		dimStyle.Render("    Find notes with specific tag") + "\n\n" +
		dimStyle.Render("  • GetAllTags() []*TagInfo") + "\n" +
		dimStyle.Render("    List all tags sorted by frequency") + "\n\n" +
		successStyle.Render("SEARCH:") + "\n" +
		successStyle.Render("PACKAGE: internal/features/search") + "\n\n" +
		codeStyle.Render("  sm := search.NewSearchManager(vaultDir)") + "\n\n" +
		dimStyle.Render("  • Search(query) ([]SearchResult, error)") + "\n" +
		dimStyle.Render("    Full-text search across all notes") + "\n\n" +
		dimStyle.Render("  • SearchInNote(path, query) ([]SearchResult, error)") + "\n" +
		dimStyle.Render("    Search within specific note") + "\n\n" +
		dimStyle.Render("  • FormatResults(results) string") + "\n" +
		dimStyle.Render("    Format results for display") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}

// GetDeveloperExamples returns code integration examples
func GetDeveloperExamples() string {
	return headerStyle.Render("💡 CODE EXAMPLES") + "\n\n" +
		successStyle.Render("COMPLETE WORKFLOW:") + "\n\n" +
		codeStyle.Render("  // Initialize managers") + "\n" +
		codeStyle.Render("  storage := storage.New()") + "\n" +
		codeStyle.Render("  statsM := stats.NewStatsManager()") + "\n" +
		codeStyle.Render("  tagM := tags.NewTagManager(vault, config)") + "\n\n" +
		codeStyle.Render("  // Create and analyze note") + "\n" +
		codeStyle.Render("  file, _ := storage.CreateNote(\"My Note\", models.FormatMarkdown)") + "\n" +
		codeStyle.Render("  content := \"# Hello\\nThis is a #test note\"") + "\n" +
		codeStyle.Render("  storage.SaveNote(file, content)") + "\n\n" +
		codeStyle.Render("  // Get statistics") + "\n" +
		codeStyle.Render("  stats := statsM.Calculate(content)") + "\n" +
		codeStyle.Render("  fmt.Printf(\"Words: %d\\n\", stats.WordCount)") + "\n\n" +
		codeStyle.Render("  // Extract and index tags") + "\n" +
		codeStyle.Render("  tags := tags.ExtractTags(content)") + "\n" +
		codeStyle.Render("  tagM.IndexNote(file.Name())") + "\n\n" +
		successStyle.Render("SEARCH & FILTER:") + "\n\n" +
		codeStyle.Render("  // Full-text search") + "\n" +
		codeStyle.Render("  searchM := search.NewSearchManager(vault)") + "\n" +
		codeStyle.Render("  results, _ := searchM.Search(\"golang\")") + "\n" +
		codeStyle.Render("  for _, r := range results {") + "\n" +
		codeStyle.Render("    fmt.Printf(\"%s: %s\\n\", r.NoteName, r.MatchSnippet)") + "\n" +
		codeStyle.Render("  }") + "\n\n" +
		codeStyle.Render("  // Find notes by tag") + "\n" +
		codeStyle.Render("  notes := tagM.GetNotesByTag(\"test\")") + "\n\n" +
		successStyle.Render("INTEGRATION TIPS:") + "\n" +
		textStyle.Render("  • All managers are thread-safe") + "\n" +
		textStyle.Render("  • Use config directory: ~/.totion/") + "\n" +
		textStyle.Render("  • JSON persistence is automatic") + "\n" +
		textStyle.Render("  • Error handling follows Go idioms") + "\n\n" +
		dimStyle.Render("Press Esc to go back")
}
