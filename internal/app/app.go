package app

import (
	"os"
	"path/filepath"
	"sync"
	"time"

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
	bridgeServer      *lingo.BridgeServer     // Bridge server manager
	selectedLangIndex int                     // Selected language index in language selector
	translating       bool                    // Track if translation is in progress
	previousState     ViewState               // Track previous state before language selector
	currentUILanguage string                  // Current UI language code (e.g., "en", "es")
	translationCache  map[string]string       // Cache for translated strings (key: "lang:text", value: "translation")
	cacheMutex        sync.RWMutex            // Mutex for thread-safe cache access
	viewCache         map[string]string       // Cache for entire rendered views (key: "lang:viewstate", value: "rendered content")
	lastCachedView    string                  // Last cached view key for quick lookup
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

	// Initialize and start Lingo.dev bridge server
	bridgeServer := lingo.NewBridgeServer()
	if lingoAPIKey != "" {
		// Only start bridge if API key is configured
		if err := bridgeServer.Start(); err != nil {
			// Log error but don't fail - app can still work without translation
			// fmt.Fprintf(os.Stderr, "Warning: Failed to start Lingo.dev bridge: %v\n", err)
		}
	}

	// Initialize Lingo.dev client with API key
	lingoClient := lingo.NewClient(lingoAPIKey)

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
		lingoClient:       lingoClient,
		bridgeServer:      bridgeServer,
		selectedLangIndex: 0,
		translating:       false,
		previousState:     ViewHome,
		currentUILanguage: "en", // Default to English,
		translationCache:  make(map[string]string),
		viewCache:         make(map[string]string),
	}

	// Setup auto-save callback
	m.autoSaver = autosave.NewAutoSaver(func() error {
		if m.isEditorDirty && m.currentFile != nil {
			if err := m.storage.SaveNote(m.currentFile, m.editor.Value()); err != nil {
				return err
			}
			m.isEditorDirty = false
			m.statusMessage = styles.SuccessStyle.Render(m.translate("✓ Auto-saved"))
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

// translate translates text to the current UI language using Lingo.dev API
// Returns original text if translation fails or language is English
// Uses caching to avoid repeated API calls - extremely fast for cached items
func (m *Model) translate(text string) string {
	// If English or no translation client, return original
	if m.currentUILanguage == "en" || m.currentUILanguage == "" {
		return text
	}

	if m.lingoClient == nil || !m.lingoClient.IsEnabled() {
		return text
	}

	// Check cache first - this is instant for cached translations
	cacheKey := m.currentUILanguage + ":" + text
	m.cacheMutex.RLock()
	cached, ok := m.translationCache[cacheKey]
	cacheSize := len(m.translationCache)
	m.cacheMutex.RUnlock()

	if ok {
		return cached
	}

	// Don't translate if we have too many items in cache (avoid memory issues)
	if cacheSize > 2000 {
		return text
	}

	// Translate using Lingo.dev API with timeout protection
	// Use quality mode (fast=false) for hackathon accuracy requirement (>90%)
	done := make(chan string, 1)
	go func() {
		// Quality mode: fast=false ensures high accuracy translation
		translated, err := m.lingoClient.TranslateText(text, "en", m.currentUILanguage, false)
		if err != nil {
			done <- text
		} else {
			done <- translated
		}
	}()

	// Wait for translation with very short timeout to prevent ANY UI freezing
	// If translation takes longer, continue in background and show result on next render
	select {
	case result := <-done:
		// Cache the result for instant future use
		m.cacheMutex.Lock()
		m.translationCache[cacheKey] = result
		m.cacheMutex.Unlock()
		return result
	case <-time.After(300 * time.Millisecond):
		// Timeout - continue translation in background, return English temporarily
		// Background goroutine will cache result for next render
		go func() {
			// Wait for background translation to complete and cache it
			result := <-done
			m.cacheMutex.Lock()
			m.translationCache[cacheKey] = result
			m.cacheMutex.Unlock()
		}()
		return text
	}
}

// clearViewCache clears the view cache when language changes
func (m *Model) clearViewCache() {
	m.viewCache = make(map[string]string)
	m.lastCachedView = ""
}

// prewarmCache pre-translates common UI strings in background for smooth rendering
// Uses batch translation API for much faster loading (translates multiple strings at once)
func (m *Model) prewarmCache(strings []string) {
	if m.currentUILanguage == "en" || m.currentUILanguage == "" {
		return
	}

	if m.lingoClient == nil || !m.lingoClient.IsEnabled() {
		return
	}

	// Filter out already-cached strings
	var untranslated []string
	m.cacheMutex.RLock()
	for _, text := range strings {
		cacheKey := m.currentUILanguage + ":" + text
		if _, ok := m.translationCache[cacheKey]; !ok {
			untranslated = append(untranslated, text)
		}
	}
	m.cacheMutex.RUnlock()

	if len(untranslated) == 0 {
		return
	}

	// Pre-translate in background using BATCH API (much faster!)
	// This translates all strings in one API call instead of sequential calls
	go func() {
		translations, err := m.lingoClient.BatchTranslateTexts(untranslated, "en", m.currentUILanguage, false)
		if err != nil {
			// Fallback to sequential if batch fails
			for _, text := range untranslated {
				translated, err := m.lingoClient.TranslateText(text, "en", m.currentUILanguage, false)
				if err == nil {
					cacheKey := m.currentUILanguage + ":" + text
					m.cacheMutex.Lock()
					m.translationCache[cacheKey] = translated
					m.cacheMutex.Unlock()
				}
			}
			return
		}

		// Cache all batch results
		m.cacheMutex.Lock()
		for original, translated := range translations {
			cacheKey := m.currentUILanguage + ":" + original
			m.translationCache[cacheKey] = translated
		}
		m.cacheMutex.Unlock()
	}()
}

// All handler functions are defined in handlers.go
// All view rendering functions are defined in views.go
