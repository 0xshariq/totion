package app

import (
	"fmt"
	"strings"

	"github.com/0xshariq/totion/internal/models"
	"github.com/0xshariq/totion/internal/notebook"
	"github.com/0xshariq/totion/internal/ui/help"
	"github.com/0xshariq/totion/internal/ui/styles"
)

// View renders the application UI
func (m *Model) View() string {
	// Focus mode - minimal UI showing only editor and word count
	if m.focusMode && m.state == ViewEditor {
		return m.renderFocusMode()
	}

	// Header with app name and tagline
	title := styles.WelcomeStyle.Render("✨ TOTION ✨")
	subtitle := styles.SubtitleStyle.Render("Your Terminal-Based Note-Taking Companion")

	header := fmt.Sprintf("%s\n%s", title, subtitle)

	var keys string
	var keysTitle string

	switch m.state {
	case ViewHome:
		keysTitle = m.translate("🎬 Quick Actions")
		keys = styles.KeysStyle.Render(
			"Ctrl+N: " + m.translate("Create New Note") + "  •  Ctrl+L: " + m.translate("View All Notes") + "  •  Ctrl+H: " + m.translate("Help") + "  •  Q: " + m.translate("Quit") + "\n" +
				"Alt+T: " + m.translate("Change UI Language") + "  •  P: " + m.translate("Themes") + "  •  S: " + m.translate("Statistics") + "  •  B: " + m.translate("Notebooks"),
		)
	case ViewList:
		keysTitle = m.translate("📋 Note List")
		keys = styles.KeysStyle.Render(
			"↑↓: " + m.translate("Navigate Notes") + "  •  Enter: " + m.translate("Open Selected") + "  •  Ctrl+D: " + m.translate("Delete Note") + "\n" +
				"/: " + m.translate("Search Notes") + "  •  Esc: " + m.translate("Back to Home"),
		)
	case ViewEditor:
		editorInfo := ""
		if m.currentNote != nil {
			pinStatus := ""
			if m.pinnedManager.IsPinned(m.currentNote.Path) {
				pinStatus = " 📌"
			}
			editorInfo = styles.StatusStyle.Render(
				fmt.Sprintf("Editing: %s %s%s", m.currentNote.Format.GetIcon(), m.currentNote.Name, pinStatus),
			)
		}
		keysTitle = m.translate("✏️  Editor Mode")
		keys = styles.KeysStyle.Render(
			"Ctrl+S: " + m.translate("Save and Close") + "  •  Alt+F: " + m.translate("Focus Mode") + "  •  Alt+P: " + m.translate("Pin/Unpin") + "\n" +
				"Alt+L: " + m.translate("Wiki Links") + "  •  Esc: " + m.translate("Minimize Editor"),
		)
		if editorInfo != "" {
			keys = editorInfo + "\n" + keys
		}
	case ViewNewFile:
		keysTitle = "📝 Create New Note"
		keys = styles.KeysStyle.Render("Enter: Proceed to Format Selection  •  Esc: Cancel & Go Back")
	case ViewFormatSelector:
		keysTitle = "🎨 Choose Format"
		keys = styles.KeysStyle.Render("Tab: Switch Between Formats  •  Enter: Create Note  •  Esc: Cancel")
	case ViewDeleteConfirm:
		keysTitle = "⚠️  Confirm Action"
		keys = styles.KeysStyle.Render("Y: Yes, Delete Permanently  •  N: No, Keep Note")
	case ViewHelp:
		keysTitle = "❓ Help & Documentation"
		keys = styles.KeysStyle.Render("Press Number: Select Topic  •  Esc: Back to Previous Menu")
	case ViewTemplates:
		keysTitle = "📝 Templates"
		keys = styles.KeysStyle.Render("1-7: Select Template Type  •  Esc: Cancel & Go Back")
	case ViewThemes:
		keysTitle = "🎨 Themes"
		keys = styles.KeysStyle.Render("1-6: Choose Theme  •  Esc: Cancel & Go Back")
	case ViewExport:
		keysTitle = "📤 Export Notes"
		keys = styles.KeysStyle.Render("1: HTML  •  2: PDF  •  3: Markdown  •  Esc: Cancel")
	case ViewImport:
		keysTitle = "📥 Import Notes"
		keys = styles.KeysStyle.Render("1: Notion  •  2: Markdown  •  3: Text Files  •  Esc: Cancel")
	case ViewLinking:
		keysTitle = "🔗 Wiki Links"
		keys = styles.KeysStyle.Render("Use [[Note Name]] syntax in editor to create links  •  Esc: Back")
	case ViewStats:
		keysTitle = "📊 Statistics Dashboard"
		keys = styles.KeysStyle.Render("View your note-taking analytics  •  Esc: Back to Home")
	case ViewGit:
		keysTitle = "🔄 Git Version Control"
		keys = styles.KeysStyle.Render("1-4: Select Git Action  •  Esc: Cancel & Go Back")
	case ViewSync:
		keysTitle = "☁️  Sync & Backup"
		keys = styles.KeysStyle.Render("1-4: Select Sync Option  •  Esc: Cancel & Go Back")
	case ViewNotebooks:
		keysTitle = "📂 Notebooks Manager"
		keys = styles.KeysStyle.Render("1-6: Select Notebook Action  •  Esc: Cancel & Go Back")
	case ViewNotebookNameInput:
		keysTitle = "📁 Create Notebook"
		keys = styles.KeysStyle.Render("Enter: Create Notebook  •  Esc: Cancel & Go Back")
	case ViewSelectNotebookForNote:
		keysTitle = "📂 Select Notebook"
		keys = styles.KeysStyle.Render("1-9: Choose Destination Notebook  •  Esc: Cancel")
	case ViewNoteNameInNotebook:
		keysTitle = "📝 Create Note in Notebook"
		keys = styles.KeysStyle.Render("Enter: Create Note  •  Esc: Cancel & Go Back")
	case ViewLanguageSelector:
		keysTitle = "🌐 UI Language Selection"
		keys = styles.KeysStyle.Render("↑↓: Navigate Languages  •  Enter: Change UI Language  •  Esc: Cancel")
	}

	var view string
	switch m.state {
	case ViewHome:
		// Show pinned notes if any
		pinnedView := m.renderPinnedNotes()

		homeContent := pinnedView +
			styles.TitleStyle.Render(m.translate("💡 WHAT IS TOTION?")) + "\n" +
			styles.InfoStyle.Render(m.translate("Totion is a powerful terminal-based note-taking application that helps you:")) + "\n" +
			styles.MenuItemStyle.Render("  • "+m.translate("Capture ideas instantly without leaving your terminal workflow")) + "\n" +
			styles.MenuItemStyle.Render("  • "+m.translate("Organize notes with notebooks, tags, and wiki-style links")) + "\n" +
			styles.MenuItemStyle.Render("  • "+m.translate("Stay focused with distraction-free focus mode")) + "\n" +
			styles.MenuItemStyle.Render("  • "+m.translate("Track your productivity with built-in analytics")) + "\n" +
			styles.MenuItemStyle.Render("  • "+m.translate("Keep everything in sync with Git integration")) + "\n\n" +

			styles.TitleStyle.Render(m.translate("🎯 QUICK START GUIDE")) + "\n" +
			styles.InfoStyle.Render(m.translate("Get started in seconds:")) + "\n" +
			styles.SuccessStyle.Render("  • Ctrl+N → "+m.translate("Create your first note")) + "\n" +
			styles.SuccessStyle.Render("  • Ctrl+L → "+m.translate("View all your notes")) + "\n" +
			styles.SuccessStyle.Render("  • Ctrl+H → "+m.translate("Open detailed help")) + "\n" +
			styles.SuccessStyle.Render("  • Q → "+m.translate("Quit application")) + "\n\n" +

			styles.TitleStyle.Render(m.translate("📝 CREATING NOTES")) + "\n" +
			styles.MenuItemStyle.Render("  • Ctrl+N → "+m.translate("Create new note with custom name")) + "\n" +
			styles.MenuItemStyle.Render("  • Alt+D → "+m.translate("Daily journal (auto-dated)")) + "\n" +
			styles.MenuItemStyle.Render("  • Ctrl+Q → "+m.translate("Quick scratch pad (temporary notes)")) + "\n" +
			styles.MenuItemStyle.Render("  • Ctrl+T → "+m.translate("Use pre-made templates (meetings, todos, etc.)")) + "\n" +
			styles.MenuItemStyle.Render("  • B → "+m.translate("Organize notes into notebooks")) + "\n" +
			styles.MenuItemStyle.Render("  • # → "+m.translate("Tag your notes for easy filtering")) + "\n" +
			styles.MenuItemStyle.Render("  • Alt+T → "+m.translate("Change UI language (English, Spanish, French, etc.)")) + "\n\n" +

			styles.TitleStyle.Render(m.translate("✏️  EDITING & FORMATTING")) + "\n" +
			styles.MenuItemStyle.Render("  • Ctrl+S → "+m.translate("Save your work and close editor")) + "\n" +
			styles.MenuItemStyle.Render("  • Alt+F → "+m.translate("Enter focus mode (minimal distractions)")) + "\n" +
			styles.MenuItemStyle.Render("  • Alt+P → "+m.translate("Pin important notes to top")) + "\n" +
			styles.MenuItemStyle.Render("  • Alt+L → "+m.translate("Create [[wiki links]] to other notes")) + "\n" +
			styles.MenuItemStyle.Render("  • P → "+m.translate("Change color themes")) + "\n" +
			styles.MenuItemStyle.Render("  • / → "+m.translate("Search within notes")) + "\n\n" +

			styles.TitleStyle.Render(m.translate("📊 VIEWING & ANALYZING")) + "\n" +
			styles.MenuItemStyle.Render("  • Ctrl+L → "+m.translate("Browse all notes in list view")) + "\n" +
			styles.MenuItemStyle.Render("  • S → "+m.translate("View statistics (note count, word count, trends)")) + "\n" +
			styles.MenuItemStyle.Render("  • Ctrl+/ → "+m.translate("Full-text search across all notes")) + "\n" +
			styles.MenuItemStyle.Render("  • ? → "+m.translate("Open help menu anytime")) + "\n\n" +

			styles.TitleStyle.Render(m.translate("💾 SYNC & BACKUP")) + "\n" +
			styles.MenuItemStyle.Render("  • Alt+E → "+m.translate("Export notes (HTML, PDF, Markdown)")) + "\n" +
			styles.MenuItemStyle.Render("  • Alt+I → "+m.translate("Import from Notion, Markdown, etc.")) + "\n" +
			styles.MenuItemStyle.Render("  • G → "+m.translate("Git integration for version control")) + "\n" +
			styles.MenuItemStyle.Render("  • Alt+Y → "+m.translate("Sync with cloud storage")) + "\n\n" +

			styles.TitleStyle.Render(m.translate("✨ WHY CHOOSE TOTION?")) + "\n" +
			styles.SuccessStyle.Render("  ✓ "+m.translate("Lightning fast - No loading times, instant startup")) + "\n" +
			styles.SuccessStyle.Render("  ✓ "+m.translate("Privacy first - All data stored locally on your machine")) + "\n" +
			styles.SuccessStyle.Render("  ✓ "+m.translate("Works offline - No internet required")) + "\n" +
			styles.SuccessStyle.Render("  ✓ "+m.translate("Auto-save - Never lose your work")) + "\n" +
			styles.SuccessStyle.Render("  ✓ "+m.translate("Keyboard driven - Maximum productivity")) + "\n" +
			styles.SuccessStyle.Render("  ✓ "+m.translate("Markdown support - Beautiful formatting")) + "\n" +
			styles.SuccessStyle.Render("  ✓ "+m.translate("Cross-platform - Works on Linux, Mac, Windows")) + "\n\n" +

			styles.WarningStyle.Render("  💡 "+m.translate("TIP: Press Ctrl+H for complete documentation and tutorials")) + "\n"

		// Set viewport content and render
		m.homeViewport.SetContent(homeContent)
		view = m.homeViewport.View() + "\n" +
			styles.ScrollHintStyle.Render(m.translate("📜 Use ↑↓ arrow keys or mouse scroll to navigate • Page Up/Down for faster scrolling"))
	case ViewList:
		view = m.list.View()
	case ViewEditor:
		view = m.editor.View()
	case ViewNewFile:
		prompt := styles.SuccessStyle.Render("Enter note name:")
		view = fmt.Sprintf("%s\n\n%s", prompt, m.fileNameInput.View())
	case ViewNotebookNameInput:
		prompt := styles.SuccessStyle.Render("📁 Create New Notebook")
		hint := styles.InfoStyle.Render("Enter a name for your notebook (e.g., Work, Personal, Projects)")
		view = fmt.Sprintf("%s\n%s\n\n%s", prompt, hint, m.notebookNameInput.View())
	case ViewSelectNotebookForNote:
		view = m.renderNotebookSelection()
	case ViewNoteNameInNotebook:
		prompt := styles.SuccessStyle.Render(fmt.Sprintf("📝 Create Note in: %s", m.selectedNotebook))
		hint := styles.InfoStyle.Render("Enter the name for your note")
		view = fmt.Sprintf("%s\n%s\n\n%s", prompt, hint, m.fileNameInput.View())
	case ViewFormatSelector:
		mdStyle := styles.StatusStyle
		txtStyle := styles.StatusStyle

		mdSelected := "  "
		txtSelected := "  "
		if m.formatIndex == 0 {
			mdSelected = "→ "
			mdStyle = styles.SuccessStyle
		} else {
			txtSelected = "→ "
			txtStyle = styles.SuccessStyle
		}

		mdLine := mdStyle.Render(fmt.Sprintf("%sMarkdown (.md) 📝", mdSelected))
		txtLine := txtStyle.Render(fmt.Sprintf("%sPlain Text (.txt) 📄", txtSelected))

		view = fmt.Sprintf("%s\n%s", mdLine, txtLine)
	case ViewDeleteConfirm:
		item, ok := m.list.SelectedItem().(models.Note)
		if ok {
			warning := styles.ErrorStyle.Render(fmt.Sprintf("⚠️  Delete %s?", item.Name))
			notice := styles.StatusStyle.Render("This action cannot be undone!")
			view = fmt.Sprintf("%s\n%s", warning, notice)
		}
	case ViewHelp:
		// Show help menu or specific topic based on selection
		var helpContent string
		if m.helpTopic == "" {
			helpContent = help.GetHelpMenu(m.translate)
		} else {
			switch m.helpTopic {
			case "1":
				helpContent = help.GetKeyboardShortcuts(m.translate)
			case "2":
				helpContent = help.GetGettingStarted(m.translate)
			case "3":
				helpContent = help.GetTemplatesGuide(m.translate)
			case "4":
				helpContent = help.GetThemesGuide(m.translate)
			case "5":
				helpContent = help.GetExportImportGuide(m.translate)
			case "6":
				helpContent = help.GetGitGuide(m.translate)
			case "7":
				helpContent = help.GetStatsGuide(m.translate)
			case "8":
				helpContent = help.GetSyncGuide(m.translate)
			case "9":
				helpContent = help.GetNotebooksGuide(m.translate)
			case "0":
				helpContent = help.GetDeveloperGuide()
			// Developer submenu (0-1 through 0-9)
			case "0-1":
				helpContent = help.GetDeveloperCorePackages()
			case "0-2":
				helpContent = help.GetDeveloperStorage()
			case "0-3":
				helpContent = help.GetDeveloperStats()
			case "0-4":
				helpContent = help.GetDeveloperExport()
			case "0-5":
				helpContent = help.GetDeveloperTemplates()
			case "0-6":
				helpContent = help.GetDeveloperAdvanced()
			case "0-7":
				helpContent = help.GetDeveloperExamples()
			case "0-8":
				helpContent = help.GetDeveloperNotebooks()
			case "0-9":
				helpContent = help.GetDeveloperBestPractices()
			case "t", "T":
				helpContent = help.GetTranslationGuide(m.translate)
			default:
				helpContent = help.GetHelpMenu(m.translate)
			}
		}
		// Set the help content in viewport and render it
		m.helpViewport.SetContent(helpContent)
		view = m.helpViewport.View() + "\n" +
			styles.ScrollHintStyle.Render(m.translate("📜 Use ↑↓ arrow keys or mouse scroll to navigate • Page Up/Down for faster scrolling"))
	case ViewTemplates:
		templates := []string{
			"1. 📋 Meeting Notes      → Structured meeting minutes",
			"2. ✅ Todo List          → Task list with checkboxes",
			"3. 📔 Journal Entry      → Daily journal with mood tracking",
			"4. 🚀 Project Plan       → Project planning template",
			"5. 💻 Code Snippet       → Save code examples",
			"6. 📚 Book Notes         → Book reading notes & highlights",
			"7. 📄 Blank Note         → Start from scratch",
		}

		templateTitle := styles.TitleStyle.Render("📝 AVAILABLE TEMPLATES")
		templateDesc := styles.InfoStyle.Render("\nSelect a template to create a new note:\n")

		var templateList string
		for i, tmpl := range templates {
			// Alternate colors for variety
			style := styles.MenuItemStyle
			if i%2 == 0 {
				style = styles.HighlightStyle
			}
			templateList += style.Render(tmpl) + "\n"
		}

		content := fmt.Sprintf("%s%s\n%s", templateTitle, templateDesc, templateList)
		m.contentViewport.SetContent(content)
		view = m.contentViewport.View() + "\n" +
			styles.ScrollHintStyle.Render("📜 Use ↑↓ arrow keys or mouse scroll if content is not fully visible")
	case ViewThemes:
		themes := []string{
			"1. 🔵 Default (Blue)     → Blue gradient, fresh and modern",
			"2. 🌙 Dark               → Blue accent, easy on the eyes",
			"3. ☀️  Light              → High contrast, bright environment",
			"4. 🎭 Monokai            → Popular developer theme",
			"5. 🌊 Solarized Dark     → Classic color scheme",
			"6. ❄️  Nord               → Cool, bluish theme",
		}

		themeTitle := styles.TitleStyle.Render("🎨 AVAILABLE THEMES")
		themeDesc := styles.InfoStyle.Render("\nSelect a theme for Totion:\n")

		var themeList string
		for i, theme := range themes {
			// Alternate colors for variety
			style := styles.MenuItemStyle
			if i%2 == 0 {
				style = styles.HighlightStyle
			}
			themeList += style.Render(theme) + "\n"
		}

		themeNote := styles.SubtleStyle.Render("\nNote: Theme selection is implemented in internal/themes")

		content := fmt.Sprintf("%s%s\n%s%s", themeTitle, themeDesc, themeList, themeNote)
		m.contentViewport.SetContent(content)
		view = m.contentViewport.View() + "\n" +
			styles.ScrollHintStyle.Render("📜 Use ↑↓ arrow keys or mouse scroll if content is not fully visible")
	case ViewExport:
		exportTitle := styles.TitleStyle.Render("📤 EXPORT OPTIONS")
		exportDesc := styles.InfoStyle.Render("\nExport your notes to different formats:\n")

		exportOptions := []string{
			"1. 📄 Export to HTML      → Beautiful web page format",
			"2. � Export to PDF       → Professional document format",
			"3. �📝 Export to Plain Text → Strip all markdown formatting",
			"4. 📋 Export to Markdown  → Preserve markdown syntax",
		}

		var exportList string
		for i, opt := range exportOptions {
			style := styles.MenuItemStyle
			if i%2 == 0 {
				style = styles.HighlightStyle
			}
			exportList += style.Render(opt) + "\n"
		}

		exportNote := styles.SubtleStyle.Render("\nPress 1-4 to select export format for current note")
		exportExample := styles.InfoStyle.Render("\nFiles are exported to /tmp/ directory")

		content := fmt.Sprintf("%s%s\n%s%s%s", exportTitle, exportDesc, exportList, exportNote, exportExample)
		m.contentViewport.SetContent(content)
		view = m.contentViewport.View() + "\n" +
			styles.ScrollHintStyle.Render("📜 Use ↑↓ arrow keys or mouse scroll if content is not fully visible")

	case ViewImport:
		importTitle := styles.TitleStyle.Render("📥 IMPORT OPTIONS")
		importDesc := styles.InfoStyle.Render("\nImport notes from other applications:\n")

		importOptions := []string{
			"1. 📦 Import from Notion     → Import Notion JSON exports",
			"2. 🔮 Import from Obsidian   → Import Obsidian vault folder",
			"3. 📁 Import Plain Text      → Batch import .txt/.md files",
		}

		var importList string
		for i, opt := range importOptions {
			style := styles.MenuItemStyle
			if i%2 == 0 {
				style = styles.HighlightStyle
			}
			importList += style.Render(opt) + "\n"
		}

		importNote := styles.SubtleStyle.Render("\nPress 1-3 to select import source")
		importExample := styles.InfoStyle.Render("\nExample: Select source folder/file to import")

		content := fmt.Sprintf("%s%s\n%s%s%s", importTitle, importDesc, importList, importNote, importExample)
		m.contentViewport.SetContent(content)
		view = m.contentViewport.View() + "\n" +
			styles.ScrollHintStyle.Render("📜 Use ↑↓ arrow keys or mouse scroll if content is not fully visible")

	case ViewLinking:
		linkingTitle := styles.TitleStyle.Render("🔗 NOTE LINKING")
		linkingDesc := styles.InfoStyle.Render("\nCreate connections between your notes:\n")

		linkingInfo := styles.MenuItemStyle.Render(`
  Wiki-Style Links:
    [[Note Title]]           → Link to another note
    [[Note|Custom Text]]     → Link with custom display text

  Features:
    • Automatic backlink tracking
    • Find all notes linking to current note
    • Build a knowledge graph
    • Connect related ideas

  Usage in Editor:
    • Type [[ to start a link
    • Write the note title
    • Close with ]]
    • Press Ctrl+K to see linking help
`)

		linkingExample := styles.HighlightStyle.Render("\nExample:")
		linkingCode := styles.CodeStyle.Render("  See [[Project Ideas]] for brainstorming\n  Check [[Meeting Notes|yesterday's meeting]]")

		view = fmt.Sprintf("%s%s%s%s\n%s", linkingTitle, linkingDesc, linkingInfo, linkingExample, linkingCode)

	case ViewStats:
		content := m.handleStatsView()
		m.contentViewport.SetContent(content)
		view = m.contentViewport.View() + "\n" +
			styles.ScrollHintStyle.Render("📜 Use ↑↓ arrow keys or mouse scroll if content is not fully visible")

	case ViewGit:
		content := m.handleGitView()
		m.contentViewport.SetContent(content)
		view = m.contentViewport.View() + "\n" +
			styles.ScrollHintStyle.Render("📜 Use ↑↓ arrow keys or mouse scroll if content is not fully visible")

	case ViewSync:
		content := m.handleSyncView()
		m.contentViewport.SetContent(content)
		view = m.contentViewport.View() + "\n" +
			styles.ScrollHintStyle.Render("📜 Use ↑↓ arrow keys or mouse scroll if content is not fully visible")

	case ViewNotebooks:
		content := m.handleNotebooksView()
		m.contentViewport.SetContent(content)
		view = m.contentViewport.View() + "\n" +
			styles.ScrollHintStyle.Render("📜 Use ↑↓ arrow keys or mouse scroll if content is not fully visible")

	case ViewLanguageSelector:
		view = m.renderLanguageSelector()
	}

	// Keyboard shortcuts section
	keysSection := ""
	if keysTitle != "" {
		keysSection = fmt.Sprintf("\n%s\n%s", styles.StatusStyle.Render(keysTitle), keys)
	}

	// Status message
	status := ""
	if m.statusMessage != "" {
		status = "\n\n" + m.statusMessage
	}

	// Build final view
	return fmt.Sprintf("\n%s\n\n%s%s%s", header, view, keysSection, status)
}

// renderNotebookSelection renders the notebook selection view
func (m *Model) renderNotebookSelection() string {
	vaultDir := m.getVaultDir()
	nbManager := notebook.NewNotebookManager(vaultDir)

	notebooks, err := nbManager.ListNotebooks()
	if err != nil || len(notebooks) == 0 {
		return styles.ErrorStyle.Render("No notebooks found. Create one first!")
	}

	title := styles.TitleStyle.Render("📂 SELECT NOTEBOOK")
	hint := styles.InfoStyle.Render("\nSelect a notebook to create your note in:\n")

	var notebookList string
	for i, nb := range notebooks {
		if i >= 9 {
			break // Only show first 9 notebooks
		}
		style := styles.MenuItemStyle
		if i%2 == 0 {
			style = styles.HighlightStyle
		}
		notebookList += style.Render(fmt.Sprintf("  %d. 📁 %s (%d notes)\n", i+1, nb.Name, nb.NoteCount))
	}

	return fmt.Sprintf("%s%s\n%s", title, hint, notebookList)
}

// renderPinnedNotes renders pinned notes section
func (m *Model) renderPinnedNotes() string {
	pinned := m.pinnedManager.GetPinned()
	if len(pinned) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(styles.TitleStyle.Render("📌 Pinned Notes:") + "\n")

	for i, p := range pinned {
		if i >= 5 {
			break // Show max 5 pinned notes
		}
		sb.WriteString(styles.HighlightStyle.Render(fmt.Sprintf("  • %s\n", p.Name)))
	}
	sb.WriteString("\n")

	return sb.String()
}

// renderFocusMode renders the minimal focus mode view
func (m *Model) renderFocusMode() string {
	var sb strings.Builder

	// Calculate stats for current editor content
	content := m.editor.Value()
	wordCount := countWords(content)
	charCount := len(content)

	// Minimal header
	sb.WriteString("\n")
	sb.WriteString(styles.SubtleStyle.Render("Focus Mode") + "\n")
	if m.currentNote != nil {
		sb.WriteString(styles.SubtleStyle.Render(m.currentNote.Name) + "\n")
	}
	sb.WriteString("\n")

	// Editor
	sb.WriteString(m.editor.View())
	sb.WriteString("\n\n")

	// Footer with stats
	stats := styles.SubtleStyle.Render(fmt.Sprintf(
		"Words: %d  •  Characters: %d  •  Ctrl+F: Exit Focus  •  Ctrl+S: Save",
		wordCount, charCount,
	))
	sb.WriteString(stats)

	return sb.String()
}

// countWords counts words in text
func countWords(text string) int {
	words := 0
	inWord := false

	for _, r := range text {
		if r == ' ' || r == '\n' || r == '\t' || r == '\r' {
			if inWord {
				words++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	if inWord {
		words++
	}

	return words
}

// renderLanguageSelector renders the language selection view
func (m *Model) renderLanguageSelector() string {
	var sb strings.Builder

	title := styles.TitleStyle.Render("🌐 SELECT UI LANGUAGE")
	subtitle := styles.InfoStyle.Render("Change the interface language to:")
	info := styles.SubtleStyle.Render("(Note: This translates menus and buttons, NOT your note content)")

	sb.WriteString(title + "\n")
	sb.WriteString(subtitle + "\n")
	sb.WriteString(info + "\n\n")

	languages := getAvailableLanguages()

	// Display in 2 columns for better layout
	for i := 0; i < len(languages); i++ {
		lang := languages[i]
		marker := "  "
		style := styles.MenuItemStyle

		if i == m.selectedLangIndex {
			marker = "→ "
			style = styles.HighlightStyle
		}

		// Single column, left-aligned
		sb.WriteString(style.Render(fmt.Sprintf("%s%s\n", marker, lang.Name)))
	}

	sb.WriteString("\n")
	sb.WriteString(styles.SubtleStyle.Render("💡 Use ↑↓ to navigate, Enter to change language, Esc to cancel"))

	return sb.String()
}
