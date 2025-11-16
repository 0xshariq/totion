package app

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/0xshariq/totion/internal/features/export"
	"github.com/0xshariq/totion/internal/features/git"
	importpkg "github.com/0xshariq/totion/internal/features/import"
	"github.com/0xshariq/totion/internal/features/linking"
	"github.com/0xshariq/totion/internal/features/stats"
	"github.com/0xshariq/totion/internal/features/sync"
	"github.com/0xshariq/totion/internal/features/templates"
	"github.com/0xshariq/totion/internal/models"
	"github.com/0xshariq/totion/internal/notebook"
	"github.com/0xshariq/totion/internal/ui/components"
	"github.com/0xshariq/totion/internal/ui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

// handleKeyPress handles keyboard input
// Returns: (handled, model, cmd)
func (m *Model) handleKeyPress(msg tea.KeyMsg) (bool, tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		if m.state == ViewHome {
			// Stop bridge server before quitting
			if m.bridgeServer != nil {
				m.bridgeServer.Stop()
			}
			// Reset language to English before quitting
			m.currentUILanguage = "en"
			m.translationCache = make(map[string]string)
			return true, m, tea.Quit
		}

	case "esc":
		// In list view with active filter, let the list handle ESC
		if m.state == ViewList && m.list.FilterState() == 1 { // 1 is Filtering state
			return false, m, nil
		}
		newModel, cmd := m.handleEscape()
		return true, newModel, cmd

	case "ctrl+l":
		newModel, cmd := m.showList()
		return true, newModel, cmd

	case "ctrl+n":
		m.state = ViewNewFile
		m.fileNameInput.SetValue("")
		m.fileNameInput.Focus()
		m.statusMessage = ""
		return true, m, nil

	case "ctrl+s":
		newModel, cmd := m.saveCurrentNote()
		return true, newModel, cmd

	case "ctrl+d":
		if m.state == ViewList {
			m.state = ViewDeleteConfirm
			m.statusMessage = m.translate("Press 'y' to confirm delete, 'n' to cancel")
			return true, m, nil
		}

	case "ctrl+h", "?":
		m.state = ViewHelp
		m.helpTopic = "" // Reset to show menu
		m.statusMessage = ""
		return true, m, nil

	case "ctrl+t":
		m.state = ViewTemplates
		m.statusMessage = ""
		return true, m, nil

	case "p", "P":
		if m.state == ViewHome {
			m.state = ViewThemes
			m.statusMessage = ""
			return true, m, nil
		}

	case "alt+e":
		if m.state == ViewHome || m.state == ViewList {
			m.state = ViewExport
			m.statusMessage = ""
			return true, m, nil
		}

	case "alt+i":
		if m.state == ViewHome {
			m.state = ViewImport
			m.statusMessage = ""
			return true, m, nil
		}

	case "alt+l":
		if m.state == ViewEditor {
			// Show linking menu in editor
			m.statusMessage = styles.InfoStyle.Render(m.translate("Use [[Note Title]] to create links"))
			return true, m, nil
		}

	case "alt+t":
		// Change UI language from anywhere in the app
		if !m.lingoClient.IsEnabled() {
			m.statusMessage = styles.WarningStyle.Render(m.translate("⚠️  Translation disabled. Set LINGODOTDEV_API_KEY in .env file. Get free key: https://lingo.dev"))
			return true, m, nil
		}
		// Save current state to return to it after language selection
		m.previousState = m.state
		m.state = ViewLanguageSelector
		m.selectedLangIndex = 0
		m.statusMessage = ""
		return true, m, nil

	case "alt+d":
		// Create or open today's daily journal (only from home view)
		if m.state == ViewHome {
			notePath, err := m.dailyManager.CreateTodayNote()
			if err != nil {
				m.statusMessage = styles.ErrorStyle.Render(m.translate("Error creating daily note: ") + err.Error())
				return true, m, nil
			}

			// Read the note content
			content, err := m.storage.ReadNote(notePath)
			if err != nil {
				m.statusMessage = styles.ErrorStyle.Render(m.translate("Error reading daily note: ") + err.Error())
				return true, m, nil
			}

			// Open the note file
			file, err := m.storage.OpenNote(notePath)
			if err != nil {
				m.statusMessage = styles.ErrorStyle.Render(m.translate("Error opening daily note: ") + err.Error())
				return true, m, nil
			}

			// Create note model
			noteInfo, err := os.Stat(notePath)
			if err != nil {
				m.statusMessage = styles.ErrorStyle.Render(m.translate("Error reading note info: ") + err.Error())
				return true, m, nil
			}

			m.currentNote = &models.Note{
				Name:    filepath.Base(notePath),
				Path:    notePath,
				Format:  models.FormatMarkdown,
				ModTime: noteInfo.ModTime(),
				Size:    noteInfo.Size(),
			}
			m.currentFile = file
			m.state = ViewEditor
			m.editor.SetValue(content)
			m.editor.Focus()

			// Add to recent notes
			if m.recentManager != nil {
				_ = m.recentManager.AddRecent(m.currentNote)
			}

			// Start auto-save
			if m.autoSaver != nil {
				m.autoSaver.Start()
			}
			m.isEditorDirty = false

			m.statusMessage = styles.SuccessStyle.Render(m.translate("📅 Daily journal opened"))
			return true, m, nil
		}

	case "alt+f":
		if m.state == ViewEditor {
			// Toggle focus mode
			m.focusMode = !m.focusMode
			if m.focusMode {
				m.statusMessage = styles.InfoStyle.Render(m.translate("Focus Mode ON - Press Ctrl+F to exit"))
			} else {
				m.statusMessage = styles.InfoStyle.Render(m.translate("Focus Mode OFF"))
			}
			return true, m, nil
		}

	case "alt+p":
		if m.state == ViewEditor && m.currentNote != nil {
			// Toggle pin for current note
			isPinned, err := m.pinnedManager.Toggle(m.currentNote.Path, m.currentNote.Name)
			if err != nil {
				m.statusMessage = styles.ErrorStyle.Render(m.translate("Error toggling pin: ") + err.Error())
			} else {
				if isPinned {
					m.statusMessage = styles.SuccessStyle.Render(m.translate("📌 Note pinned!"))
				} else {
					m.statusMessage = styles.InfoStyle.Render(m.translate("📌 Note unpinned"))
				}
			}
			return true, m, nil
		}

	case "s", "S":
		if m.state == ViewHome {
			m.state = ViewStats
			m.statusMessage = ""
			return true, m, nil
		}

	case "g", "G":
		if m.state == ViewHome {
			m.state = ViewGit
			m.statusMessage = ""
			return true, m, nil
		}

	case "alt+y":
		if m.state == ViewHome {
			m.state = ViewSync
			m.statusMessage = ""
			return true, m, nil
		}

	case "b", "B":
		if m.state == ViewHome {
			m.state = ViewNotebooks
			m.statusMessage = ""
			return true, m, nil
		}

	case "enter":
		if m.state == ViewLanguageSelector {
			return m.translateNote()
		}
		if m.state == ViewNewFile || m.state == ViewFormatSelector || m.state == ViewList || m.state == ViewNotebookNameInput || m.state == ViewNoteNameInNotebook {
			newModel, cmd := m.handleEnter()
			return true, newModel, cmd
		}

	case "up", "k":
		if m.state == ViewLanguageSelector && m.selectedLangIndex > 0 {
			m.selectedLangIndex--
			return true, m, nil
		}

	case "down", "j":
		if m.state == ViewLanguageSelector {
			languages := getAvailableLanguages()
			if m.selectedLangIndex < len(languages)-1 {
				m.selectedLangIndex++
			}
			return true, m, nil
		}

	case "tab":
		if m.state == ViewFormatSelector {
			m.formatIndex = (m.formatIndex + 1) % 2
			if m.formatIndex == 0 {
				m.selectedFormat = models.FormatMarkdown
			} else {
				m.selectedFormat = models.FormatText
			}
			return true, m, nil
		}

	case "y", "Y":
		if m.state == ViewDeleteConfirm {
			newModel, cmd := m.deleteSelectedNote()
			return true, newModel, cmd
		}

	case "n", "N":
		if m.state == ViewDeleteConfirm {
			m.state = ViewList
			m.statusMessage = ""
			return true, m, nil
		}

	case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
		if m.state == ViewHelp && m.helpTopic == "" {
			// Select main help topic
			m.helpTopic = msg.String()
			return true, m, nil
		}
		// Developer submenu uses 1-9
		if m.state == ViewHelp && m.helpTopic == "0" {
			m.helpTopic = "0-" + msg.String() // e.g., "0-1", "0-2", etc.
			return true, m, nil
		}
		// Only handle numeric input for non-help views
		if m.state == ViewTemplates {
			newModel, cmd := m.selectTemplate(msg.String())
			return true, newModel, cmd
		}
		if m.state == ViewThemes {
			m.selectTheme(msg.String())
			return true, m, nil
		}
		if m.state == ViewExport {
			m.handleExport(msg.String())
			return true, m, nil
		}
		if m.state == ViewImport {
			m.handleImport(msg.String())
			return true, m, nil
		}
		if m.state == ViewGit {
			m.handleGitAction(msg.String())
			return true, m, nil
		}
		if m.state == ViewSync {
			m.handleSyncAction(msg.String())
			return true, m, nil
		}
		if m.state == ViewNotebooks {
			m.handleNotebookAction(msg.String())
			return true, m, nil
		}
		if m.state == ViewSelectNotebookForNote {
			m.selectNotebookForNote(msg.String())
			return true, m, nil
		}

	case "t", "T":
		if m.state == ViewHome {
			m.state = ViewTags
			m.statusMessage = ""
			return true, m, nil
		}
	}

	// Key not handled globally, let component handle it
	return false, m, nil
}

// handleEscape handles the escape key
func (m *Model) handleEscape() (tea.Model, tea.Cmd) {
	switch m.state {
	case ViewHelp:
		// If viewing a developer subsection (0-1, 0-2, etc), go back to developer menu
		if len(m.helpTopic) > 2 && m.helpTopic[0] == '0' && m.helpTopic[1] == '-' {
			m.helpTopic = "0" // Back to developer menu
		} else if m.helpTopic != "" {
			m.helpTopic = "" // Back to main help menu
		} else {
			m.state = ViewHome
		}
		m.statusMessage = ""

	case ViewNewFile, ViewFormatSelector, ViewTemplates, ViewThemes, ViewExport, ViewImport, ViewLinking, ViewStats, ViewGit, ViewSync, ViewNotebooks, ViewNotebookNameInput, ViewSelectNotebookForNote, ViewNoteNameInNotebook, ViewLanguageSelector:
		// Clear inputs before going home
		m.notebookNameInput.SetValue("")
		m.fileNameInput.SetValue("")
		m.selectedNotebook = ""
		m.selectedLangIndex = 0
		m.translating = false
		m.state = ViewHome
		m.statusMessage = ""

	case ViewDeleteConfirm:
		m.state = ViewList
		m.statusMessage = ""

	case ViewEditor:
		// Just go back to home without saving
		// Don't close file or stop auto-save - let auto-save continue in background
		m.state = ViewHome
		m.statusMessage = styles.InfoStyle.Render(m.translate("Editor minimized - auto-save still active. Press Ctrl+L to view notes"))

	case ViewList:
		m.state = ViewHome
		m.statusMessage = ""
	}

	return m, nil
}

// handleEnter handles the enter key
func (m *Model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.state {
	case ViewNewFile:
		filename := m.fileNameInput.Value()
		if filename != "" {
			m.state = ViewFormatSelector
			m.formatIndex = 0
			m.selectedFormat = models.FormatMarkdown
			m.statusMessage = ""
		}
		return m, nil

	case ViewFormatSelector:
		// Check if we're creating a note in a notebook
		if m.selectedNotebook != "" {
			return m.createNoteInNotebook()
		}
		return m.createNewNote()

	case ViewList:
		return m.openSelectedNote()

	case ViewNotebookNameInput:
		return m.createNotebook()

	case ViewNoteNameInNotebook:
		// Move to format selector
		filename := m.fileNameInput.Value()
		if filename != "" {
			m.state = ViewFormatSelector
			m.formatIndex = 0
			m.selectedFormat = models.FormatMarkdown
			m.statusMessage = ""
		}
		return m, nil
	}

	return m, nil
}

// showList shows the note list
func (m *Model) showList() (tea.Model, tea.Cmd) {
	notes, err := m.storage.ListNotes()
	if err != nil {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Error: ") + err.Error())
		return m, nil
	}

	m.list = components.NewNoteList(notes)
	h, v := styles.DocStyle.GetFrameSize()
	m.list.SetSize(m.width-h, m.height-v-5)
	m.state = ViewList
	m.statusMessage = ""

	return m, nil
}

// createNewNote creates a new note
func (m *Model) createNewNote() (tea.Model, tea.Cmd) {
	filename := m.fileNameInput.Value()

	file, path, err := m.storage.CreateNote(filename, m.selectedFormat)
	if err != nil {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Error: ") + err.Error())
		m.state = ViewHome
		return m, nil
	}

	// Get template content if a template was selected
	templateContent := ""
	if m.selectedTemplate != "" {
		tm := templates.NewTemplateManager()
		template, err := tm.GetTemplate(m.selectedTemplate)
		if err == nil {
			templateContent = template.Content
		}
		// Reset selected template
		m.selectedTemplate = ""
	}

	m.currentFile = file
	m.currentNote = &models.Note{
		Name:   filename + m.selectedFormat.GetExtension(),
		Path:   path,
		Format: m.selectedFormat,
	}
	m.state = ViewEditor
	m.editor.SetValue(templateContent)
	m.editor.Focus()
	m.statusMessage = styles.SuccessStyle.Render(fmt.Sprintf(m.translate("Created %s %s"), m.selectedFormat.GetIcon(), filename))

	return m, nil
}

// openSelectedNote opens the selected note from the list
func (m *Model) openSelectedNote() (tea.Model, tea.Cmd) {
	if m.currentFile != nil {
		return m, nil
	}

	item, ok := m.list.SelectedItem().(models.Note)
	if !ok {
		return m, nil
	}

	content, err := m.storage.ReadNote(item.Path)
	if err != nil {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Error: ") + err.Error())
		return m, nil
	}

	file, err := m.storage.OpenNote(item.Path)
	if err != nil {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Error: ") + err.Error())
		return m, nil
	}

	m.editor.SetValue(content)
	m.currentFile = file
	m.currentNote = &item
	m.state = ViewEditor

	// Add to recent notes
	if m.recentManager != nil {
		_ = m.recentManager.AddRecent(&item)
	}

	// Start auto-save
	if m.autoSaver != nil {
		m.autoSaver.Start()
	}
	m.isEditorDirty = false

	// Parse and show links if any
	linker := linking.NewLinkManager()
	links := linker.ParseLinks(content, item.Name)
	if len(links) > 0 {
		linkInfo := m.translate(fmt.Sprintf("Found %d wiki-style links. Press Alt+L for link help.", len(links)))
		m.statusMessage = styles.InfoStyle.Render(linkInfo)
	} else {
		m.statusMessage = styles.StatusStyle.Render(fmt.Sprintf(m.translate("Editing %s %s"), item.Format.GetIcon(), item.Name))
	}

	return m, nil
}

// saveCurrentNote saves the current note
func (m *Model) saveCurrentNote() (tea.Model, tea.Cmd) {
	if m.currentFile == nil {
		return m, nil
	}

	err := m.storage.SaveNote(m.currentFile, m.editor.Value())
	if err != nil {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Error: ") + err.Error())
		return m, nil
	}

	// Close the file handle
	if m.currentFile != nil {
		m.currentFile.Close()
	}

	// Stop auto-save
	if m.autoSaver != nil {
		m.autoSaver.Stop()
	}
	m.isEditorDirty = false

	m.currentFile = nil
	m.currentNote = nil
	m.editor.SetValue("")
	m.state = ViewHome
	m.statusMessage = styles.SuccessStyle.Render(m.translate("✓ Note saved successfully!"))

	return m, nil
}

// deleteSelectedNote deletes the selected note
func (m *Model) deleteSelectedNote() (tea.Model, tea.Cmd) {
	item, ok := m.list.SelectedItem().(models.Note)
	if !ok {
		m.state = ViewList
		return m, nil
	}

	err := m.storage.DeleteNote(item.Path)
	if err != nil {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Error: " + err.Error()))
		m.state = ViewList
		return m, nil
	}

	m.statusMessage = styles.SuccessStyle.Render(fmt.Sprintf(m.translate("✓ Deleted %s"), item.Name))
	m.state = ViewHome

	return m, nil
}

// selectTemplate selects and applies a template
func (m *Model) selectTemplate(key string) (tea.Model, tea.Cmd) {
	templateNames := []string{
		m.translate("Meeting Notes"),
		m.translate("Todo List"),
		m.translate("Journal Entry"),
		m.translate("Project Plan"),
		m.translate("Code Snippet"),
		m.translate("Book Notes"),
		m.translate("Blank"),
	}

	index := 0
	switch key {
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	case "4":
		index = 3
	case "5":
		index = 4
	case "6":
		index = 5
	case "7":
		index = 6
	default:
		return m, nil
	}

	// Store selected template name
	m.selectedTemplate = templateNames[index]

	// Transition to new file creation with template
	m.state = ViewNewFile
	m.selectedFormat = models.FormatMarkdown
	m.fileNameInput.SetValue("")
	m.fileNameInput.Focus()
	m.statusMessage = styles.SuccessStyle.Render(m.translate(fmt.Sprintf("Selected: %s template", templateNames[index])))

	return m, nil
}

// selectTheme selects and applies a theme
func (m *Model) selectTheme(key string) {
	themeNames := []string{
		"Default (Blue)",
		"Dark",
		"Light",
		"Monokai",
		"Solarized Dark",
		"Nord",
	}

	index := 0
	switch key {
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	case "4":
		index = 3
	case "5":
		index = 4
	case "6":
		index = 5
	default:
		return
	}

	m.state = ViewHome
	m.statusMessage = styles.SuccessStyle.Render(fmt.Sprintf(m.translate("✓ Theme changed to: %s"), themeNames[index]))
}

// handleExport handles export operations
func (m *Model) handleExport(key string) {
	if m.currentNote == nil {
		m.state = ViewHome
		m.statusMessage = styles.ErrorStyle.Render(m.translate("⚠️  No note open. Open a note first to export."))
		return
	}

	exporter := export.NewExporter()
	content := m.editor.Value()
	outputPath := "/tmp/" + m.currentNote.Name

	var err error
	var format string

	switch key {
	case "1": // HTML
		err = exporter.ExportToHTML(content, m.currentNote.Name, outputPath+".html")
		format = "HTML"
	case "2": // PDF
		err = exporter.ExportToPDF(content, m.currentNote.Name, outputPath+".pdf")
		format = "PDF"
	case "3": // Plain Text
		err = exporter.ExportToPlainText(content, outputPath+".txt")
		format = "Plain Text"
	case "4": // Markdown
		err = exporter.ExportToMarkdown(content, outputPath+".md")
		format = "Markdown"
	default:
		return
	}

	if err != nil {
		m.statusMessage = styles.ErrorStyle.Render(fmt.Sprintf(m.translate("Export failed: %v"), err))
	} else {
		m.statusMessage = styles.SuccessStyle.Render(fmt.Sprintf(m.translate("✓ Exported to %s format at %s"), format, outputPath))
	}
	m.state = ViewHome
}

// handleImport handles import operations
func (m *Model) handleImport(key string) {
	// Get vault directory from storage
	vaultDir := os.ExpandEnv("$HOME/.totion")
	importer := importpkg.NewImporter(vaultDir)
	var message string

	switch key {
	case "1": // Notion
		// Check if file exists
		notionPath := filepath.Join(os.Getenv("HOME"), "notion_export.json")
		if _, err := os.Stat(notionPath); os.IsNotExist(err) {
			message = fmt.Sprintf("Place notion_export.json in %s and try again", os.Getenv("HOME"))
			m.statusMessage = styles.WarningStyle.Render(m.translate("⚠️  ") + message)
		} else {
			imported, err := importer.ImportFromNotion(notionPath)
			if err != nil {
				m.statusMessage = styles.ErrorStyle.Render(fmt.Sprintf(m.translate("Import failed: %v"), err))
			} else {
				m.statusMessage = styles.SuccessStyle.Render(fmt.Sprintf(m.translate("✓ Imported %d notes from Notion"), len(imported)))
			}
		}
	case "2": // Obsidian
		obsidianPath := filepath.Join(os.Getenv("HOME"), "obsidian_vault")
		if _, err := os.Stat(obsidianPath); os.IsNotExist(err) {
			message = fmt.Sprintf("Place obsidian_vault/ folder in %s and try again", os.Getenv("HOME"))
			m.statusMessage = styles.WarningStyle.Render(m.translate("⚠️  ") + message)
		} else {
			imported, err := importer.ImportFromObsidian(obsidianPath)
			if err != nil {
				m.statusMessage = styles.ErrorStyle.Render(fmt.Sprintf(m.translate("Import failed: %v"), err))
			} else {
				m.statusMessage = styles.SuccessStyle.Render(fmt.Sprintf(m.translate("✓ Imported %d notes from Obsidian"), len(imported)))
			}
		}
	case "3": // Plain Text
		// For plain text, show instructions
		message = "Place .md or .txt files in ~/.totion/ manually, or use git clone to import from repositories"
		m.statusMessage = styles.InfoStyle.Render(m.translate("💡 ") + message)
	default:
		return
	}

	m.state = ViewHome
}

// getVaultDir returns the vault directory path
func (m *Model) getVaultDir() string {
	return filepath.Join(os.Getenv("HOME"), ".totion")
}

// handleGitAction performs git actions
func (m *Model) handleGitAction(key string) {
	vaultDir := m.getVaultDir()
	gitManager := git.NewGitManager(vaultDir)

	switch key {
	case "1": // Initialize
		err := gitManager.Initialize()
		if err != nil {
			m.statusMessage = styles.ErrorStyle.Render(fmt.Sprintf(m.translate("Git init failed: %v"), err))
		} else {
			m.statusMessage = styles.SuccessStyle.Render(m.translate("✓ Git repository initialized"))
		}
	case "2": // Commit
		err := gitManager.Commit("Auto-commit from Totion")
		if err != nil {
			m.statusMessage = styles.ErrorStyle.Render(fmt.Sprintf(m.translate("Commit failed: %v"), err))
		} else {
			m.statusMessage = styles.SuccessStyle.Render(m.translate("✓ Changes committed"))
		}
	case "3": // History
		history, _ := gitManager.GetHistory(5)
		m.statusMessage = styles.InfoStyle.Render(fmt.Sprintf(m.translate("Last 5 commits:\n%s"), history))
		return
	case "4": // Status
		status, _ := gitManager.GetStatus()
		m.statusMessage = styles.InfoStyle.Render(fmt.Sprintf(m.translate("Git status:\n%s"), status))
		return
	default:
		return
	}
	m.state = ViewHome
}

// handleSyncAction performs sync/backup actions
func (m *Model) handleSyncAction(key string) {
	vaultDir := m.getVaultDir()
	syncDir := filepath.Join(os.TempDir(), "totion_sync")
	syncer := sync.NewSyncManager(vaultDir, syncDir)

	switch key {
	case "1": // Backup
		backupPath := filepath.Join(os.TempDir(), "totion_backup")
		err := syncer.BackupVault(backupPath)
		if err != nil {
			m.statusMessage = styles.ErrorStyle.Render(fmt.Sprintf(m.translate("Backup failed: %v"), err))
		} else {
			m.statusMessage = styles.SuccessStyle.Render(fmt.Sprintf(m.translate("✓ Backup created: %s"), backupPath))
		}
	case "2": // Restore
		backupPath := filepath.Join(os.TempDir(), "totion_backup")
		err := syncer.RestoreVault(backupPath)
		if err != nil {
			m.statusMessage = styles.ErrorStyle.Render(fmt.Sprintf(m.translate("Restore failed: %v"), err))
		} else {
			m.statusMessage = styles.SuccessStyle.Render(m.translate("✓ Vault restored from backup"))
		}
	case "3": // Sync to cloud
		m.statusMessage = styles.InfoStyle.Render(m.translate("Cloud sync requires configuration. See README for setup."))
		return
	case "4": // Sync from cloud
		m.statusMessage = styles.InfoStyle.Render(m.translate("Cloud sync requires configuration. See README for setup."))
		return
	default:
		return
	}
	m.state = ViewHome
}

// handleStatsView shows statistics for the current note or vault
func (m *Model) handleStatsView() string {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".totion")
	statsManager := stats.NewStatsManagerWithConfig(configDir)
	vaultDir := m.getVaultDir()
	nbManager := notebook.NewNotebookManager(vaultDir)

	// Get all notes for vault stats
	notes, _ := m.storage.ListNotes()
	notebooks, _ := nbManager.ListNotebooks()

	var totalWords, totalChars int

	for _, note := range notes {
		content, err := m.storage.ReadNote(note.Path)
		if err == nil {
			noteStats := statsManager.Calculate(content)
			totalWords += noteStats.WordCount
			totalChars += noteStats.CharCount
		}
	}

	// Record today's stats
	statsManager.RecordStats(time.Now(), totalWords, len(notes))

	// Build notebook stats
	notebookStats := []stats.NotebookStats{}
	for _, nb := range notebooks {
		notebookStats = append(notebookStats, stats.NotebookStats{
			Name:      nb.Name,
			NoteCount: nb.NoteCount,
			WordCount: 0, // Could calculate if needed
		})
	}

	// Sort by note count
	sort.Slice(notebookStats, func(i, j int) bool {
		return notebookStats[i].NoteCount > notebookStats[j].NoteCount
	})

	// Build dashboard data
	dashboardData := statsManager.BuildDashboard(len(notes), len(notebooks), notebookStats)

	// Render beautiful dashboard
	return stats.RenderDashboard(dashboardData)
}

// handleGitView shows git operations
func (m *Model) handleGitView() string {
	vaultDir := m.getVaultDir()
	gitManager := git.NewGitManager(vaultDir)

	gitTitle := styles.TitleStyle.Render(m.translate("🔄 GIT VERSION CONTROL"))

	// Check if git is initialized
	status, _ := gitManager.GetStatus()

	gitInfo := fmt.Sprintf(`
%s: %s

%s:
  1. %s
  2. %s
  3. %s
  4. %s

%s:
%s

`,
		m.translate("Git Repository"), vaultDir,
		m.translate("Options"),
		m.translate("Initialize Git Repository"),
		m.translate("Commit All Changes"),
		m.translate("View Commit History"),
		m.translate("View Status"),
		m.translate("Status"), status)

	return gitTitle + "\n" + styles.MenuItemStyle.Render(gitInfo)
}

// handleSyncView shows sync/backup options
func (m *Model) handleSyncView() string {
	syncTitle := styles.TitleStyle.Render(m.translate("☁️  SYNC & BACKUP"))

	syncInfo := fmt.Sprintf(`
%s:
  1. %s
  2. %s
  3. %s
  4. %s

%s: ~/.totion/

%s
`,
		m.translate("Sync Options"),
		m.translate("Backup Vault to ZIP"),
		m.translate("Restore from Backup"),
		m.translate("Sync to Cloud (if configured)"),
		m.translate("Sync from Cloud"),
		m.translate("Current vault location"),
		m.translate("Press 1-4 to select an option"))

	return syncTitle + "\n" + styles.MenuItemStyle.Render(syncInfo)
}

// handleNotebooksView shows notebook management
func (m *Model) handleNotebooksView() string {
	vaultDir := m.getVaultDir()
	nbManager := notebook.NewNotebookManager(vaultDir)

	notebooksTitle := styles.TitleStyle.Render(m.translate("📂 NOTEBOOKS & FOLDERS"))

	notebooksInfo := fmt.Sprintf(`
%s:
  1. %s
  2. %s
  3. %s
  4. %s
  5. %s
  6. %s

%s
%s

%s
`,
		m.translate("Notebook Management"),
		m.translate("Create New Notebook"),
		m.translate("List All Notebooks"),
		m.translate("Move Note to Notebook"),
		m.translate("Rename Notebook"),
		m.translate("Delete Notebook"),
		m.translate("Create Note in Notebook"),
		m.translate("Notebooks help organize your notes into folders."),
		m.translate("Each notebook is a folder in ~/.totion/"),
		m.translate("Press 1-6 to select an option"))

	// List existing notebooks (directories)
	notebooks, _ := nbManager.ListNotebooks()
	notebookList := ""
	if len(notebooks) > 0 {
		notebookList = styles.HighlightStyle.Render(m.translate("\nExisting Notebooks:\n"))
		for _, nb := range notebooks {
			notebookList += styles.InfoStyle.Render(fmt.Sprintf("  📁 %s\n", nb.Name))
		}
	}

	return notebooksTitle + "\n" + styles.MenuItemStyle.Render(notebooksInfo) + notebookList
}

// handleNotebookAction performs notebook actions
func (m *Model) handleNotebookAction(key string) {
	vaultDir := m.getVaultDir()
	nbManager := notebook.NewNotebookManager(vaultDir)

	switch key {
	case "1": // Create New Notebook
		// Switch to notebook name input state
		m.notebookNameInput.SetValue("")
		m.notebookNameInput.Placeholder = "Enter notebook name..."
		m.notebookNameInput.Focus()
		m.state = ViewNotebookNameInput
		m.statusMessage = styles.InfoStyle.Render(m.translate("Enter notebook name and press Enter"))
	case "2": // List All Notebooks
		notebooks, err := nbManager.ListNotebooks()
		if err != nil || len(notebooks) == 0 {
			m.statusMessage = styles.InfoStyle.Render(m.translate("No notebooks found. Create one with option 1."))
		} else {
			notebookNames := ""
			for _, nb := range notebooks {
				notebookNames += fmt.Sprintf("  📁 %s (%d notes)\n", nb.Name, nb.NoteCount)
			}
			m.statusMessage = styles.InfoStyle.Render(fmt.Sprintf(m.translate("Notebooks:\n%s"), notebookNames))
		}
		m.state = ViewHome
	case "3": // Move Note to Notebook
		m.statusMessage = styles.InfoStyle.Render(m.translate("Feature coming soon: Select note and destination notebook"))
		m.state = ViewHome
	case "4": // Rename Notebook
		m.statusMessage = styles.InfoStyle.Render(m.translate("Feature coming soon: Select notebook to rename"))
		m.state = ViewHome
	case "5": // Delete Notebook
		m.statusMessage = styles.InfoStyle.Render(m.translate("Feature coming soon: Select notebook to delete"))
		m.state = ViewHome
	case "6": // Create Note in Notebook
		// Show list of notebooks to select from
		notebooks, err := nbManager.ListNotebooks()
		if err != nil || len(notebooks) == 0 {
			m.statusMessage = styles.ErrorStyle.Render(m.translate("No notebooks found. Create one first with option 1."))
			m.state = ViewHome
		} else {
			m.state = ViewSelectNotebookForNote
			m.statusMessage = styles.InfoStyle.Render(m.translate("Select a notebook (type the number):"))
		}
	default:
		return
	}
}

// createNotebook creates a new notebook folder
func (m *Model) createNotebook() (tea.Model, tea.Cmd) {
	notebookName := m.notebookNameInput.Value()

	if notebookName == "" {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Notebook name cannot be empty"))
		return m, nil
	}

	vaultDir := m.getVaultDir()
	nbManager := notebook.NewNotebookManager(vaultDir)

	err := nbManager.CreateNotebook(notebookName)
	if err != nil {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Error creating notebook: ") + err.Error())
		m.state = ViewHome
		return m, nil
	}

	m.statusMessage = styles.SuccessStyle.Render(fmt.Sprintf(m.translate("✓ Created notebook: %s"), notebookName))
	m.state = ViewHome
	m.notebookNameInput.SetValue("")

	return m, nil
}

// selectNotebookForNote handles notebook selection for creating a note
func (m *Model) selectNotebookForNote(key string) {
	vaultDir := m.getVaultDir()
	nbManager := notebook.NewNotebookManager(vaultDir)

	notebooks, err := nbManager.ListNotebooks()
	if err != nil || len(notebooks) == 0 {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("No notebooks available"))
		m.state = ViewHome
		return
	}

	// Convert key to index
	index := -1
	switch key {
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	case "4":
		index = 3
	case "5":
		index = 4
	case "6":
		index = 5
	case "7":
		index = 6
	case "8":
		index = 7
	case "9":
		index = 8
	}

	if index >= 0 && index < len(notebooks) {
		m.selectedNotebook = notebooks[index].Name
		m.state = ViewNoteNameInNotebook
		m.fileNameInput.SetValue("")
		m.fileNameInput.Focus()
		m.statusMessage = styles.InfoStyle.Render(fmt.Sprintf(m.translate("Creating note in: %s"), m.selectedNotebook))
	}
}

// createNoteInNotebook creates a note in the selected notebook
func (m *Model) createNoteInNotebook() (tea.Model, tea.Cmd) {
	filename := m.fileNameInput.Value()

	if filename == "" {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Note name cannot be empty"))
		return m, nil
	}

	if m.selectedNotebook == "" {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("No notebook selected"))
		m.state = ViewHome
		return m, nil
	}

	vaultDir := m.getVaultDir()
	notebookPath := filepath.Join(vaultDir, m.selectedNotebook)

	// Create note file in the notebook folder
	ext := m.selectedFormat.GetExtension()
	notePath := filepath.Join(notebookPath, filename+ext)

	// Check if file already exists
	if _, err := os.Stat(notePath); err == nil {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Note already exists in this notebook"))
		return m, nil
	}

	file, err := os.Create(notePath)
	if err != nil {
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Error creating note: ") + err.Error())
		m.state = ViewHome
		return m, nil
	}

	// Get template content if a template was selected
	templateContent := ""
	if m.selectedTemplate != "" {
		tm := templates.NewTemplateManager()
		template, err := tm.GetTemplate(m.selectedTemplate)
		if err == nil {
			templateContent = template.Content
		}
		m.selectedTemplate = ""
	}

	m.currentFile = file
	m.currentNote = &models.Note{
		Name:   filename + ext,
		Path:   notePath,
		Format: m.selectedFormat,
	}
	m.state = ViewEditor
	m.editor.SetValue(templateContent)
	m.editor.Focus()
	m.statusMessage = styles.SuccessStyle.Render(fmt.Sprintf(m.translate("Created %s in %s"), filename+ext, m.selectedNotebook))

	// Reset selected notebook
	m.selectedNotebook = ""

	return m, nil
}

// Language represents a supported translation language
type Language struct {
	Code string
	Name string
}

// getAvailableLanguages returns the list of supported translation languages
func getAvailableLanguages() []Language {
	return []Language{
		{"es", "ES  Spanish (Español)"},
		{"fr", "FR  French (Français)"},
		{"de", "DE  German (Deutsch)"},
		{"it", "IT  Italian (Italiano)"},
		{"pt", "PT  Portuguese (Português)"},
		{"ja", "JP  Japanese (日本語)"},
		{"ko", "KR  Korean (한국어)"},
		{"zh-Hans", "CN  Chinese Simplified (简体中文)"},
		{"zh-Hant", "TW  Chinese Traditional (繁體中文)"},
		{"ru", "RU  Russian (Русский)"},
		{"ar", "SA  Arabic (العربية)"},
		{"hi", "IN  Hindi (हिन्दी)"},
		{"nl", "NL  Dutch (Nederlands)"},
		{"pl", "PL  Polish (Polski)"},
		{"tr", "TR  Turkish (Türkçe)"},
		{"sv", "SE  Swedish (Svenska)"},
		{"no", "NO  Norwegian (Norsk)"},
		{"da", "DK  Danish (Dansk)"},
		{"fi", "FI  Finnish (Suomi)"},
		{"el", "GR  Greek (Ελληνικά)"},
		{"cs", "CZ  Czech (Čeština)"},
		{"ro", "RO  Romanian (Română)"},
		{"hu", "HU  Hungarian (Magyar)"},
		{"vi", "VN  Vietnamese (Tiếng Việt)"},
		{"th", "TH  Thai (ไทย)"},
		{"id", "ID  Indonesian (Bahasa Indonesia)"},
		{"he", "IL  Hebrew (עברית)"},
		{"uk", "UA  Ukrainian (Українська)"},
		{"ca", "ES  Catalan (Català)"},
		{"bg", "BG  Bulgarian (Български)"},
		{"hr", "HR  Croatian (Hrvatski)"},
		{"sk", "SK  Slovak (Slovenčina)"},
		{"sl", "SI  Slovenian (Slovenščina)"},
		{"lt", "LT  Lithuanian (Lietuvių)"},
		{"lv", "LV  Latvian (Latviešu)"},
		{"et", "EE  Estonian (Eesti)"},
		{"ms", "MY  Malay (Bahasa Melayu)"},
		{"tl", "PH  Filipino (Filipino)"},
	}
}

// translateNote switches the UI language and translates all UI elements
func (m *Model) translateNote() (bool, tea.Model, tea.Cmd) {
	languages := getAvailableLanguages()
	if m.selectedLangIndex >= len(languages) {
		m.state = m.previousState
		m.statusMessage = styles.ErrorStyle.Render(m.translate("Invalid language selection"))
		return true, m, nil
	}

	targetLang := languages[m.selectedLangIndex]

	// Update the current UI language
	m.currentUILanguage = targetLang.Code

	// Clear translation cache when language changes
	m.translationCache = make(map[string]string)
	m.clearViewCache()

	// Pre-warm cache with common strings for smooth rendering (100+ strings)
	commonStrings := []string{
		// Main actions
		"Create New Note", "View All Notes", "Help", "Quit",
		"Change UI Language", "Themes", "Statistics", "Notebooks",
		// Navigation
		"Navigate Notes", "Open Selected", "Delete Note", "Search Notes",
		"Back to Home", "Save and Close", "Focus Mode", "Pin/Unpin",
		"Wiki Links", "Minimize Editor", "Quick Actions", "Note List",
		"Editor Mode",
		// Status messages
		"Note saved successfully!", "Note pinned!", "Note unpinned",
		"Daily journal opened", "Focus Mode ON - Press Ctrl+F to exit",
		"Focus Mode OFF",
		// Common UI elements
		"Enter note name:", "Create New Notebook", "Editing",
		"Press Esc to go back", "Use [[Note Title]] to create links",
		// View titles
		"🎬 Quick Actions", "📋 Note List", "✏️  Editor Mode",
		"📝 Create New Note", "❓ Help & Documentation",
		// Keyboard shortcuts
		"Ctrl+N", "Ctrl+L", "Ctrl+H", "Ctrl+S", "Alt+T", "Alt+F",
		// Common words
		"Cancel", "Confirm", "Yes", "No", "Error", "Success",
		// Home page titles
		"💡 WHAT IS TOTION?", "🎯 QUICK START GUIDE", "📝 CREATING NOTES",
		"✏️  EDITING & FORMATTING", "📊 VIEWING & ANALYZING", "💾 SYNC & BACKUP",
		"✨ WHY CHOOSE TOTION?",
		// Home page descriptions
		"Totion is a powerful terminal-based note-taking application that helps you:",
		"Get started in seconds:",
		// Home page bullet points - WHAT IS TOTION
		"Capture ideas instantly without leaving your terminal workflow",
		"Organize notes with notebooks, tags, and wiki-style links",
		"Stay focused with distraction-free focus mode",
		"Track your productivity with built-in analytics",
		"Keep everything in sync with Git integration",
		// QUICK START GUIDE bullets
		"Create your first note", "View all your notes", "Open detailed help", "Quit application",
		// CREATING NOTES bullets
		"Create new note with custom name", "Daily journal (auto-dated)",
		"Quick scratch pad (temporary notes)", "Use pre-made templates (meetings, todos, etc.)",
		"Organize notes into notebooks", "Tag your notes for easy filtering",
		"Change UI language (English, Spanish, French, etc.)",
		// EDITING & FORMATTING bullets
		"Save your work and close editor", "Enter focus mode (minimal distractions)",
		"Pin important notes to top", "Create [[wiki links]] to other notes",
		"Change color themes", "Search within notes",
		// VIEWING & ANALYZING bullets
		"Browse all notes in list view", "View statistics (note count, word count, trends)",
		"Full-text search across all notes", "Open help menu anytime",
		// SYNC & BACKUP bullets
		"Export notes (HTML, PDF, Markdown)", "Import from Notion, Markdown, etc.",
		"Git integration for version control", "Sync with cloud storage",
		// WHY CHOOSE TOTION bullets
		"Lightning fast - No loading times, instant startup",
		"Privacy first - All data stored locally on your machine",
		"Works offline - No internet required",
		"Auto-save - Never lose your work",
		"Keyboard driven - Maximum productivity",
		"Markdown support - Beautiful formatting",
		"Cross-platform - Works on Linux, Mac, Windows",
		"TIP: Press Ctrl+H for complete documentation and tutorials",
		// Scroll hints
		"📜 Use ↑↓ arrow keys or mouse scroll to navigate • Page Up/Down for faster scrolling",
		"📜 Use ↑↓ arrow keys or mouse scroll if content is not fully visible",
	}
	m.prewarmCache(commonStrings)

	// Return to previous state
	m.state = m.previousState

	// Show appropriate status message with translation stats
	if targetLang.Code == "en" {
		m.statusMessage = styles.SuccessStyle.Render(m.translate("✓ UI language set to English"))
	} else if m.lingoClient == nil || !m.lingoClient.IsEnabled() {
		m.statusMessage = styles.WarningStyle.Render(
			fmt.Sprintf(m.translate("⚠ UI set to %s but translation unavailable - set LINGODOTDEV_API_KEY in .env file"), targetLang.Name))
	} else {
		// Count cached translations
		m.cacheMutex.RLock()
		cacheSize := len(m.translationCache)
		m.cacheMutex.RUnlock()

		m.statusMessage = styles.SuccessStyle.Render(
			fmt.Sprintf(m.translate("✓ UI language set to %s - %d strings translated"), targetLang.Name, cacheSize))
	}

	return true, m, nil
}
