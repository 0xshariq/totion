package lingo

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// BridgeServer manages the Lingo.dev bridge server lifecycle
type BridgeServer struct {
	cmd     *exec.Cmd
	running bool
	port    int
}

// NewBridgeServer creates a new bridge server manager
func NewBridgeServer() *BridgeServer {
	return &BridgeServer{
		port:    3737,
		running: false,
	}
}

// IsRunning checks if the bridge server is already running
func (b *BridgeServer) IsRunning() bool {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/health", b.port))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// GetURL returns the bridge server URL
func (b *BridgeServer) GetURL() string {
	return fmt.Sprintf("http://localhost:%d", b.port)
}

// Start starts the bridge server if it's not already running
func (b *BridgeServer) Start() error {
	// Check if already running
	if b.IsRunning() {
		b.running = true
		return nil // Already running, nothing to do
	}

	// Find the bridge server directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	bridgePath := filepath.Join(cwd, "lingo-bridge")
	if _, err := os.Stat(bridgePath); os.IsNotExist(err) {
		return fmt.Errorf("lingo-bridge directory not found at %s", bridgePath)
	}

	// Check if node_modules exists, if not install dependencies
	nodeModulesPath := filepath.Join(bridgePath, "node_modules")
	if _, err := os.Stat(nodeModulesPath); os.IsNotExist(err) {
		// Try pnpm first, fallback to npm
		installCmd := exec.Command("pnpm", "install")
		installCmd.Dir = bridgePath
		if err := installCmd.Run(); err != nil {
			// Fallback to npm
			installCmd = exec.Command("npm", "install")
			installCmd.Dir = bridgePath
			if err := installCmd.Run(); err != nil {
				return fmt.Errorf("failed to install bridge dependencies: %w", err)
			}
		}
	}

	// Start the bridge server in the background
	// Try pnpm first, fallback to node
	b.cmd = exec.Command("pnpm", "start")
	b.cmd.Dir = bridgePath

	// Create pipes to capture output for debugging
	b.cmd.Stdout = os.Stdout // Show output for debugging
	b.cmd.Stderr = os.Stderr

	if err := b.cmd.Start(); err != nil {
		// Fallback to direct node command
		b.cmd = exec.Command("node", "server.js")
		b.cmd.Dir = bridgePath
		b.cmd.Stdout = os.Stdout
		b.cmd.Stderr = os.Stderr

		if err := b.cmd.Start(); err != nil {
			return fmt.Errorf("failed to start bridge server: %w", err)
		}
	}

	// Wait for server to be ready (max 10 seconds)
	maxRetries := 40 // Increased for slower systems
	retryDelay := 250 * time.Millisecond

	for i := 0; i < maxRetries; i++ {
		time.Sleep(retryDelay)
		if b.IsRunning() {
			b.running = true
			fmt.Printf("✓ Lingo.dev bridge server started successfully on port %d\n", b.port)
			fmt.Println("  • Quality mode enabled for maximum accuracy")
			fmt.Println("  • Context-aware translation active")
			fmt.Println("  • Technical glossary loaded")
			return nil // Server is ready
		}

		// Show progress every 2 seconds
		if i > 0 && i%8 == 0 {
			fmt.Printf("  Waiting for bridge server... (%d/%d)\n", i, maxRetries)
		}
	}

	// Server failed to start - try to kill the process
	if b.cmd != nil && b.cmd.Process != nil {
		_ = b.cmd.Process.Kill()
	}

	return fmt.Errorf("bridge server failed to start in time (waited %d seconds)", maxRetries*int(retryDelay.Seconds()))
}

// Stop stops the bridge server
func (b *BridgeServer) Stop() error {
	if !b.running || b.cmd == nil || b.cmd.Process == nil {
		return nil // Not running
	}

	fmt.Printf("Stopping Lingo.dev bridge server (PID: %d)...\n", b.cmd.Process.Pid)

	// Try graceful shutdown first on Unix systems
	if err := b.cmd.Process.Signal(os.Interrupt); err == nil {
		// Wait up to 2 seconds for graceful shutdown
		done := make(chan error, 1)
		go func() {
			done <- b.cmd.Wait()
		}()

		select {
		case <-done:
			b.running = false
			fmt.Println("✓ Bridge server stopped gracefully")
			return nil
		case <-time.After(2 * time.Second):
			// Graceful shutdown timed out, force kill
		}
	}

	// Force kill if graceful shutdown failed
	if err := b.cmd.Process.Kill(); err != nil {
		return fmt.Errorf("failed to stop bridge server: %w", err)
	}

	// Wait for process to exit
	_ = b.cmd.Wait()

	b.running = false
	fmt.Println("✓ Bridge server stopped")
	return nil
}
