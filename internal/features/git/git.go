package git

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

// GitManager handles git operations for version control
type GitManager struct {
	vaultDir string
}

// NewGitManager creates a new git manager
func NewGitManager(vaultDir string) *GitManager {
	return &GitManager{
		vaultDir: vaultDir,
	}
}

// Initialize initializes a git repository
func (gm *GitManager) Initialize() error {
	cmd := exec.Command("git", "init", gm.vaultDir)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error initializing git: %w", err)
	}

	// Create .gitignore
	gitignorePath := filepath.Join(gm.vaultDir, ".gitignore")
	gitignoreContent := []byte("*.tmp\n.DS_Store\n")

	if err := exec.Command("sh", "-c", fmt.Sprintf("echo '%s' > %s", gitignoreContent, gitignorePath)).Run(); err != nil {
		return fmt.Errorf("error creating .gitignore: %w", err)
	}

	return nil
}

// Commit commits changes with a message
func (gm *GitManager) Commit(message string) error {
	// Add all files
	addCmd := exec.Command("git", "-C", gm.vaultDir, "add", ".")
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("error staging files: %w", err)
	}

	// Commit
	commitCmd := exec.Command("git", "-C", gm.vaultDir, "commit", "-m", message)
	if err := commitCmd.Run(); err != nil {
		return fmt.Errorf("error committing: %w", err)
	}

	return nil
}

// GetStatus gets git status
func (gm *GitManager) GetStatus() (string, error) {
	cmd := exec.Command("git", "-C", gm.vaultDir, "status", "--short")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting status: %w", err)
	}

	return string(output), nil
}

// GetHistory gets commit history
func (gm *GitManager) GetHistory(limit int) (string, error) {
	cmd := exec.Command("git", "-C", gm.vaultDir, "log", fmt.Sprintf("-%d", limit), "--oneline")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting history: %w", err)
	}

	return string(output), nil
}

// IsRepository checks if vault is a git repository
func (gm *GitManager) IsRepository() bool {
	cmd := exec.Command("git", "-C", gm.vaultDir, "rev-parse", "--git-dir")
	return cmd.Run() == nil
}

// AutoCommit automatically commits changes
func (gm *GitManager) AutoCommit(filename string) error {
	message := fmt.Sprintf("Auto-save: %s", filename)
	return gm.Commit(message)
}
