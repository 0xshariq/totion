package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

// EncryptionManager handles note encryption/decryption
type EncryptionManager struct{}

// NewEncryptionManager creates a new encryption manager
func NewEncryptionManager() *EncryptionManager {
	return &EncryptionManager{}
}

// Encrypt encrypts content with a password
func (em *EncryptionManager) Encrypt(content, password string) (string, error) {
	// Derive key from password
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("error generating salt: %w", err)
	}

	key := pbkdf2.Key([]byte(password), salt, 10000, 32, sha256.New)

	// Create cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("error creating cipher: %w", err)
	}

	// Encrypt data
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creating GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("error generating nonce: %w", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(content), nil)

	// Combine salt and ciphertext
	result := append(salt, ciphertext...)

	return base64.StdEncoding.EncodeToString(result), nil
}

// Decrypt decrypts content with a password
func (em *EncryptionManager) Decrypt(encrypted, password string) (string, error) {
	// Decode base64
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("error decoding data: %w", err)
	}

	if len(data) < 32 {
		return "", fmt.Errorf("invalid encrypted data")
	}

	// Extract salt and ciphertext
	salt := data[:32]
	ciphertext := data[32:]

	// Derive key
	key := pbkdf2.Key([]byte(password), salt, 10000, 32, sha256.New)

	// Create cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("error creating cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creating GCM: %w", err)
	}

	if len(ciphertext) < gcm.NonceSize() {
		return "", fmt.Errorf("invalid ciphertext")
	}

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]

	// Decrypt
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("error decrypting: %w", err)
	}

	return string(plaintext), nil
}

// IsEncrypted checks if content appears to be encrypted
func (em *EncryptionManager) IsEncrypted(content string) bool {
	// Simple check: encrypted content should be base64
	_, err := base64.StdEncoding.DecodeString(content)
	return err == nil && len(content) > 44 // Min encrypted data size
}
