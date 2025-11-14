package app

import (
	"os"
	"path/filepath"

	"github.com/0xshariq/totion/internal/features/autosave"
	"github.com/0xshariq/totion/internal/features/daily"
	"github.com/0xshariq/totion/internal/features/pinned"
	"github.com/0xshariq/totion/internal/features/quick"
	"github.com/0xshariq/totion/internal/features/recent"
	"github.com/0xshariq/totion/internal/features/search"
	"github.com/0xshariq/totion/internal/features/tags"
	"github.com/0xshariq/totion/internal/lingo"
	"github.com/0xshariq/totion/internal/models"
	"github.com/0xshariq/totion/internal/storage"
	"github.com/0xshariq/totion/internal/ui/components"
	"github.com/0xshariq/totion/internal/ui/styles"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

// ViewState represents the current view state
type ViewState int

const (
	ViewHome ViewState = iota
	ViewList
	ViewEditor
	ViewNewFile
	ViewFormatSelector
	ViewDeleteConfirm
	ViewHelp
	ViewTemplates
	ViewThemes
	ViewExport
	ViewImport
	ViewLinking
	ViewStats
	ViewGit
	ViewSync
	ViewNotebooks
	ViewNotebookNameInput
	ViewSelectNotebookForNote
	ViewNoteNameInNotebook
	ViewDailyNote
	ViewQuickNote
	ViewSearch
	ViewTags
	ViewLanguageSelector
)

// Model represents the main application model
type Model struct {
	storage           *storage.Storage
	state             ViewState
	list              list.Model
	editor            textarea.Model
	fileNameInput     textinput.Model
	notebookNameInput textinput.Model // Input for notebook names
	searchInput       textinput.Model // Input for search queries
	homeViewport      viewport.Model  // Viewport for scrolling home view
	helpViewport      viewport.Model  // Viewport for scrolling help content
	contentViewport   viewport.Model  // General viewport for other scrollable views
	currentFile       *os.File
	currentNote       *models.Note
	selectedFormat    models.FileFormat
	selectedTemplate  string // Track selected template name
	selectedNotebook  string // Track selected notebook for note creation
	helpTopic         string // Track selected help topic
	formatIndex       int
	statusMessage     string
	err               error
	width             int
	height            int
	autoSaver         *autosave.AutoSaver     // Auto-save manager
	recentManager     *recent.RecentManager   // Recent notes manager
	pinnedManager     *pinned.PinnedManager   // Pinned notes manager
	dailyManager      *daily.DailyManager     // Daily notes manager
	quickManager      *quick.QuickNoteManager // Quick note manager
	searchManager     *search.SearchManager   // Search manager
	tagManager        *tags.TagManager        // Tag manager
	searchResults     []search.SearchResult   // Search results
	isEditorDirty     bool                    // Track if editor has unsaved changes
	focusMode         bool                    // Focus mode (minimal UI)
	homeViewReady     bool                    // Track if home viewport is initialized
	lingoClient       *lingo.Client           // Lingo.dev translation client
	selectedLangIndex int                     // Selected language index in language selector
	translating       bool                    // Track if translation is in progress
	previousState     ViewState               // Track previous state before language selector
	currentUILanguage string                  // Current UI language code (e.g., "en", "es")
}

// New creates a new application model
func New() *Model {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".totion")

	// Load .env file from current directory first, then from home directory
	_ = godotenv.Load(".env")                           // Try current directory
	_ = godotenv.Load(filepath.Join(homeDir, ".env"))   // Try home directory
	_ = godotenv.Load(filepath.Join(configDir, ".env")) // Try config directory

	// Get API key from .env file, fallback to environment variable
	lingoAPIKey := os.Getenv("LINGODOTDEV_API_KEY")
	if lingoAPIKey == "" {
		// Fallback: check for alternative env var names
		lingoAPIKey = os.Getenv("LINGO_API_KEY")
	}

	// Initialize Lingo.dev client
	lingoURL := os.Getenv("LINGO_BRIDGE_URL")
	if lingoURL == "" {
		lingoURL = "http://localhost:3000"
	}

	m := &Model{
		storage:           storage.New(),
		state:             ViewHome,
		editor:            components.NewEditor(),
		fileNameInput:     components.NewFileNameInput(),
		notebookNameInput: components.NewFileNameInput(),
		selectedFormat:    models.FormatMarkdown,
		formatIndex:       0,
		recentManager:     recent.NewRecentManager(configDir),
		pinnedManager:     pinned.NewPinnedManager(configDir),
		focusMode:         false,
		lingoClient:       lingo.NewClient(lingoURL),
		selectedLangIndex: 0,
		translating:       false,
		previousState:     ViewHome,
		currentUILanguage: "en", // Default to English
	}

	// Setup auto-save callback
	m.autoSaver = autosave.NewAutoSaver(func() error {
		if m.isEditorDirty && m.currentFile != nil {
			if err := m.storage.SaveNote(m.currentFile, m.editor.Value()); err != nil {
				return err
			}
			m.isEditorDirty = false
			m.statusMessage = styles.SuccessStyle.Render("✓ Auto-saved")
		}
		return nil
	})

	return m
}

// Init initializes the model
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Initialize home viewport if not ready
		if !m.homeViewReady {
			m.homeViewport = viewport.New(msg.Width, msg.Height-5) // -5 for status bar
			m.homeViewport.YPosition = 0
			m.homeViewReady = true
		} else {
			m.homeViewport.Width = msg.Width
			m.homeViewport.Height = msg.Height - 5
		}

		// Initialize/update help viewport
		m.helpViewport = viewport.New(msg.Width, msg.Height-10) // -10 for header and keys
		m.helpViewport.YPosition = 0

		// Initialize/update general content viewport for other views
		m.contentViewport = viewport.New(msg.Width, msg.Height-10) // -10 for header and keys
		m.contentViewport.YPosition = 0

		// Only set list size if list is initialized (in ViewList state)
		if m.state == ViewList {
			h, v := styles.DocStyle.GetFrameSize()
			m.list.SetSize(msg.Width-h, msg.Height-v-5)
		}

	case tea.KeyMsg:
		// Handle global shortcuts first
		handled, newModel, newCmd := m.handleKeyPress(msg)
		if handled {
			return newModel, newCmd
		}
		// If not handled globally, pass to active component
	}

	// Update active component based on state
	switch m.state {
	case ViewHome:
		// Allow viewport scrolling in home view
		m.homeViewport, cmd = m.homeViewport.Update(msg)
	case ViewHelp:
		// Allow viewport scrolling in help view
		m.helpViewport, cmd = m.helpViewport.Update(msg)
	case ViewTemplates, ViewThemes, ViewExport, ViewImport, ViewStats, ViewGit, ViewSync, ViewNotebooks:
		// Allow viewport scrolling in content views
		m.contentViewport, cmd = m.contentViewport.Update(msg)
	case ViewList:
		m.list, cmd = m.list.Update(msg)
	case ViewEditor:
		m.editor, cmd = m.editor.Update(msg)
	case ViewNewFile, ViewNoteNameInNotebook:
		m.fileNameInput, cmd = m.fileNameInput.Update(msg)
	case ViewNotebookNameInput:
		m.notebookNameInput, cmd = m.notebookNameInput.Update(msg)
	}

	return m, cmd
}

// All handler functions are defined in handlers.go
// All view rendering functions are defined in views.go
