package sync

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// SyncManager handles cloud sync operations
type SyncManager struct {
	vaultDir string
	syncDir  string
	enabled  bool
}

// NewSyncManager creates a new sync manager
func NewSyncManager(vaultDir, syncDir string) *SyncManager {
	return &SyncManager{
		vaultDir: vaultDir,
		syncDir:  syncDir,
		enabled:  false,
	}
}

// Enable enables sync
func (sm *SyncManager) Enable() {
	sm.enabled = true
}

// Disable disables sync
func (sm *SyncManager) Disable() {
	sm.enabled = false
}

// IsEnabled checks if sync is enabled
func (sm *SyncManager) IsEnabled() bool {
	return sm.enabled
}

// SyncToCloud syncs local notes to cloud storage
func (sm *SyncManager) SyncToCloud() error {
	if !sm.enabled {
		return fmt.Errorf("sync is not enabled")
	}

	if err := os.MkdirAll(sm.syncDir, 0750); err != nil {
		return fmt.Errorf("error creating sync directory: %w", err)
	}

	return filepath.Walk(sm.vaultDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(sm.vaultDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(sm.syncDir, relPath)

		if err := os.MkdirAll(filepath.Dir(destPath), 0750); err != nil {
			return err
		}

		return copyFile(path, destPath)
	})
}

// SyncFromCloud syncs cloud storage to local notes
func (sm *SyncManager) SyncFromCloud() error {
	if !sm.enabled {
		return fmt.Errorf("sync is not enabled")
	}

	return filepath.Walk(sm.syncDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(sm.syncDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(sm.vaultDir, relPath)

		if err := os.MkdirAll(filepath.Dir(destPath), 0750); err != nil {
			return err
		}

		return copyFile(path, destPath)
	})
}

// BackupVault creates a backup of the entire vault
func (sm *SyncManager) BackupVault(backupPath string) error {
	return filepath.Walk(sm.vaultDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(sm.vaultDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(backupPath, relPath)

		if err := os.MkdirAll(filepath.Dir(destPath), 0750); err != nil {
			return err
		}

		return copyFile(path, destPath)
	})
}

// RestoreVault restores vault from backup
func (sm *SyncManager) RestoreVault(backupPath string) error {
	return filepath.Walk(backupPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(backupPath, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(sm.vaultDir, relPath)

		if err := os.MkdirAll(filepath.Dir(destPath), 0750); err != nil {
			return err
		}

		return copyFile(path, destPath)
	})
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
