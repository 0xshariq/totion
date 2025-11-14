package main

import (
	"fmt"
	"log"
	"os"

	"github.com/0xshariq/totion/internal/app"
	"github.com/0xshariq/totion/internal/config"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Initialize configuration
	if err := config.Initialize(); err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	// Create and run the application
	model := app.New()
	p := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running application: %v\n", err)
		os.Exit(1)
	}
}
