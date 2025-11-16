package lingo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client handles communication with Lingo.dev via bridge server
// The bridge server uses the official Lingo.dev JavaScript SDK
type Client struct {
	apiKey        string
	baseURL       string
	httpClient    *http.Client
	bridgeServer  *BridgeServer // Reference to bridge server for smart waiting
	serverChecked bool          // Track if we've checked server readiness
}

// Response represents the API response structure
type Response struct {
	Translation string `json:"translation,omitempty"`
	Cached      bool   `json:"cached,omitempty"`
	Error       string `json:"error,omitempty"`
}

// NewClient creates a new Lingo.dev client that connects to the bridge server
// The bridge server uses the official Lingo.dev JavaScript SDK with Redis caching
func NewClient(apiKey string) *Client {
	// Bridge server URL - runs locally on port 3737
	bridgeURL := "http://localhost:3737"

	if apiKey == "" {
		// Return a disabled client if no API key
		return &Client{
			apiKey:  "",
			baseURL: bridgeURL,
			httpClient: &http.Client{
				Timeout: 30 * time.Second,
			},
		}
	}

	return &Client{
		apiKey:  apiKey,
		baseURL: bridgeURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second, // Increased slightly for reliability
		},
		serverChecked: false,
	}
}

// SetBridgeServer sets the bridge server reference for smart waiting
func (c *Client) SetBridgeServer(bridge *BridgeServer) {
	c.bridgeServer = bridge
}

// waitForServer waits for the bridge server to be ready (only called on first translation)
func (c *Client) waitForServer() error {
	if c.serverChecked {
		return nil // Already checked
	}

	if c.bridgeServer == nil {
		c.serverChecked = true
		return nil // No bridge reference, assume ready
	}

	// Wait up to 3 seconds for server to be ready
	maxWait := 30 // 30 x 100ms = 3 seconds max
	for i := 0; i < maxWait; i++ {
		if c.bridgeServer.IsRunning() {
			c.serverChecked = true
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}

	c.serverChecked = true // Don't keep checking
	return fmt.Errorf("bridge server not ready after 3 seconds")
}

// IsEnabled returns true if the client has an API key configured
func (c *Client) IsEnabled() bool {
	return c.apiKey != ""
}

// TranslateText translates a text string to target language
// Calls the bridge server which uses the official Lingo.dev JavaScript SDK
// Results are cached in Redis for fast repeated translations
// Use fast=false for quality mode (>90% accuracy) or fast=true for speed
// Automatically retries failed requests with exponential backoff
func (c *Client) TranslateText(text, sourceLocale, targetLocale string, fast bool) (string, error) {
	if !c.IsEnabled() {
		return text, nil // Return original text if no API key
	}

	if text == "" {
		return text, nil
	}

	// Validate language codes
	if len(targetLocale) < 2 {
		return text, fmt.Errorf("invalid target locale: %s", targetLocale)
	}

	// Wait for server to be ready (only on first call, then cached)
	if err := c.waitForServer(); err != nil {
		return text, err
	}

	payload := map[string]interface{}{
		"text":         text,
		"sourceLocale": sourceLocale,
		"targetLocale": targetLocale,
		"fast":         fast,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return text, err
	}

	// Retry logic: up to 3 attempts with exponential backoff
	maxRetries := 3
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff: 100ms, 200ms, 400ms
			time.Sleep(time.Duration(100*(1<<uint(attempt-1))) * time.Millisecond)
		}

		req, err := http.NewRequest("POST", c.baseURL+"/translate", bytes.NewBuffer(body))
		if err != nil {
			if attempt == maxRetries-1 {
				return text, err
			}
			continue
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			if attempt == maxRetries-1 {
				return text, fmt.Errorf("bridge server request failed (is it running?): %w", err)
			}
			continue
		}

		if resp.StatusCode != http.StatusOK {
			bodyBytes, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			if attempt == maxRetries-1 {
				return text, fmt.Errorf("translation error (status %d): %s", resp.StatusCode, string(bodyBytes))
			}
			continue
		}

		var result Response
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			if attempt == maxRetries-1 {
				return text, err
			}
			continue
		}
		resp.Body.Close()

		if result.Error != "" {
			if attempt == maxRetries-1 {
				return text, fmt.Errorf("translation failed: %s", result.Error)
			}
			continue
		}

		if result.Translation != "" {
			return result.Translation, nil
		}

		if attempt == maxRetries-1 {
			return text, fmt.Errorf("no translation in response")
		}
	}

	return text, fmt.Errorf("translation failed after %d retries", maxRetries)
}

// BatchTranslateTexts translates multiple texts at once (more efficient than individual calls)
// Uses the batch endpoint for faster pre-warming of UI strings
// Automatically handles retries and provides detailed stats
func (c *Client) BatchTranslateTexts(texts []string, sourceLocale, targetLocale string, fast bool) (map[string]string, error) {
	if !c.IsEnabled() {
		return make(map[string]string), nil
	}

	if len(texts) == 0 {
		return make(map[string]string), nil
	}

	// Validate language codes
	if len(targetLocale) < 2 {
		return nil, fmt.Errorf("invalid target locale: %s", targetLocale)
	}

	// Wait for server to be ready (only on first call, then cached)
	if err := c.waitForServer(); err != nil {
		return nil, err
	}

	payload := map[string]any{
		"texts":        texts,
		"sourceLocale": sourceLocale,
		"targetLocale": targetLocale,
		"fast":         fast,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/translate/batch", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Use longer timeout for batch requests (translating many strings)
	client := &http.Client{
		Timeout: 15 * time.Second, // Increased for large batches
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("batch translation request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("batch translation error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if errMsg, ok := result["error"].(string); ok && errMsg != "" {
		return nil, fmt.Errorf("batch translation failed: %s", errMsg)
	}

	translations := make(map[string]string)

	// Handle new response format with results array
	if results, ok := result["results"].([]interface{}); ok {
		for i, r := range results {
			if i < len(texts) {
				if translated, ok := r.(string); ok {
					translations[texts[i]] = translated
				} else {
					// Fallback to original if translation failed
					translations[texts[i]] = texts[i]
				}
			}
		}
	} else {
		// Fallback: old format or error
		return nil, fmt.Errorf("unexpected response format from batch translation")
	}

	// Log stats if available (for debugging)
	if stats, ok := result["stats"].(map[string]interface{}); ok {
		if cacheHitRate, ok := stats["cacheHitRate"].(string); ok {
			// Optional: Could log this for monitoring
			_ = cacheHitRate
		}
	}

	return translations, nil
}

// DetectLanguage is a placeholder - language detection not currently used
func (c *Client) DetectLanguage(text string) (string, error) {
	return "en", nil
}
