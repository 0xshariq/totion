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

// Start starts the bridge server if it's not already running (DEPRECATED - use StartAsync)
func (b *BridgeServer) Start() error {
	return b.StartAsync()
}

// StartAsync starts the bridge server asynchronously - returns immediately
// The server starts in the background and becomes available within 1-2 seconds
func (b *BridgeServer) StartAsync() error {
	// Quick check if already running (fast, non-blocking)
	if b.IsRunning() {
		b.running = true
		return nil
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
		// Install in background - don't block
		go func() {
			// Try pnpm first, fallback to npm
			installCmd := exec.Command("pnpm", "install")
			installCmd.Dir = bridgePath
			installCmd.Stdout = nil
			installCmd.Stderr = nil
			if err := installCmd.Run(); err != nil {
				// Fallback to npm
				installCmd = exec.Command("npm", "install")
				installCmd.Dir = bridgePath
				installCmd.Stdout = nil
				installCmd.Stderr = nil
				_ = installCmd.Run()
			}
		}()
	}

	// Start the bridge server in the background
	// Try node first (faster startup than pnpm)
	b.cmd = exec.Command("node", "server.js")
	b.cmd.Dir = bridgePath
	b.cmd.Stdout = nil
	b.cmd.Stderr = nil

	if err := b.cmd.Start(); err != nil {
		// Fallback to pnpm
		b.cmd = exec.Command("pnpm", "start")
		b.cmd.Dir = bridgePath
		b.cmd.Stdout = nil
		b.cmd.Stderr = nil
		if err := b.cmd.Start(); err != nil {
			return fmt.Errorf("failed to start bridge server: %w", err)
		}
	}

	// Mark as starting (don't wait for it to be ready)
	b.running = true

	// Return immediately - server will be ready in ~1-2 seconds
	// Translation requests will automatically wait for server to be ready
	return nil
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
