package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	VaultDir      string
	DefaultFormat string
}

var AppConfig *Config

// Initialize sets up the application configuration
func Initialize() error {
	// Load .env file from current directory or parent directories
	// godotenv.Load() will silently fail if .env doesn't exist, which is fine
	_ = godotenv.Load()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	vaultDir := filepath.Join(homeDir, ".totion")

	// Create vault directory if it doesn't exist
	if err := os.MkdirAll(vaultDir, 0750); err != nil {
		return fmt.Errorf("error creating vault directory: %w", err)
	}

	AppConfig = &Config{
		VaultDir:      vaultDir,
		DefaultFormat: "md",
	}

	return nil
}
