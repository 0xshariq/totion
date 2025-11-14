package themes

import "github.com/charmbracelet/lipgloss"

// Theme represents a color theme
type Theme struct {
	Name       string
	Primary    lipgloss.Color
	Secondary  lipgloss.Color
	Background lipgloss.Color
	Foreground lipgloss.Color
	Accent     lipgloss.Color
	Error      lipgloss.Color
	Success    lipgloss.Color
}

// ThemeManager handles theme management
type ThemeManager struct {
	themes       map[string]Theme
	currentTheme string
}

// NewThemeManager creates a new theme manager
func NewThemeManager() *ThemeManager {
	tm := &ThemeManager{
		themes:       make(map[string]Theme),
		currentTheme: "default",
	}

	tm.initializeThemes()
	return tm
}

// initializeThemes sets up built-in themes
func (tm *ThemeManager) initializeThemes() {
	// Default (Pink) Theme
	tm.themes["default"] = Theme{
		Name:       "Default",
		Primary:    lipgloss.Color("205"),
		Secondary:  lipgloss.Color("241"),
		Background: lipgloss.Color("16"),
		Foreground: lipgloss.Color("254"),
		Accent:     lipgloss.Color("212"),
		Error:      lipgloss.Color("196"),
		Success:    lipgloss.Color("46"),
	}

	// Dark Theme
	tm.themes["dark"] = Theme{
		Name:       "Dark",
		Primary:    lipgloss.Color("39"),
		Secondary:  lipgloss.Color("245"),
		Background: lipgloss.Color("235"),
		Foreground: lipgloss.Color("255"),
		Accent:     lipgloss.Color("33"),
		Error:      lipgloss.Color("196"),
		Success:    lipgloss.Color("46"),
	}

	// Light Theme
	tm.themes["light"] = Theme{
		Name:       "Light",
		Primary:    lipgloss.Color("27"),
		Secondary:  lipgloss.Color("240"),
		Background: lipgloss.Color("255"),
		Foreground: lipgloss.Color("16"),
		Accent:     lipgloss.Color("33"),
		Error:      lipgloss.Color("160"),
		Success:    lipgloss.Color("28"),
	}

	// Monokai Theme
	tm.themes["monokai"] = Theme{
		Name:       "Monokai",
		Primary:    lipgloss.Color("197"),
		Secondary:  lipgloss.Color("245"),
		Background: lipgloss.Color("234"),
		Foreground: lipgloss.Color("230"),
		Accent:     lipgloss.Color("81"),
		Error:      lipgloss.Color("197"),
		Success:    lipgloss.Color("148"),
	}

	// Solarized Dark Theme
	tm.themes["solarized-dark"] = Theme{
		Name:       "Solarized Dark",
		Primary:    lipgloss.Color("33"),
		Secondary:  lipgloss.Color("240"),
		Background: lipgloss.Color("234"),
		Foreground: lipgloss.Color("244"),
		Accent:     lipgloss.Color("37"),
		Error:      lipgloss.Color("160"),
		Success:    lipgloss.Color("64"),
	}

	// Nord Theme
	tm.themes["nord"] = Theme{
		Name:       "Nord",
		Primary:    lipgloss.Color("81"),
		Secondary:  lipgloss.Color("245"),
		Background: lipgloss.Color("235"),
		Foreground: lipgloss.Color("252"),
		Accent:     lipgloss.Color("109"),
		Error:      lipgloss.Color("203"),
		Success:    lipgloss.Color("108"),
	}
}

// GetTheme returns a theme by name
func (tm *ThemeManager) GetTheme(name string) (Theme, bool) {
	theme, ok := tm.themes[name]
	return theme, ok
}

// GetCurrentTheme returns the current theme
func (tm *ThemeManager) GetCurrentTheme() Theme {
	return tm.themes[tm.currentTheme]
}

// SetTheme sets the current theme
func (tm *ThemeManager) SetTheme(name string) bool {
	if _, ok := tm.themes[name]; ok {
		tm.currentTheme = name
		return true
	}
	return false
}

// ListThemes returns all available theme names
func (tm *ThemeManager) ListThemes() []string {
	names := make([]string, 0, len(tm.themes))
	for name := range tm.themes {
		names = append(names, name)
	}
	return names
}

// AddCustomTheme adds a custom theme
func (tm *ThemeManager) AddCustomTheme(theme Theme) {
	tm.themes[theme.Name] = theme
}
