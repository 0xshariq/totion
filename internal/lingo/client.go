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
	apiKey     string
	baseURL    string
	httpClient *http.Client
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
			Timeout: 2 * time.Second, // Short timeout to prevent UI freezing
		},
	}
}

// IsEnabled returns true if the client has an API key configured
func (c *Client) IsEnabled() bool {
	return c.apiKey != ""
}

// TranslateText translates a text string to target language
// Calls the bridge server which uses the official Lingo.dev JavaScript SDK
// Results are cached in Redis for fast repeated translations
// Use fast=false for quality mode (>90% accuracy) or fast=true for speed
func (c *Client) TranslateText(text, sourceLocale, targetLocale string, fast bool) (string, error) {
	if !c.IsEnabled() {
		return text, nil // Return original text if no API key
	}

	if text == "" {
		return text, nil
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

	req, err := http.NewRequest("POST", c.baseURL+"/translate", bytes.NewBuffer(body))
	if err != nil {
		return text, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return text, fmt.Errorf("bridge server request failed (is it running?): %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return text, fmt.Errorf("translation error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return text, err
	}

	if result.Error != "" {
		return text, fmt.Errorf("translation failed: %s", result.Error)
	}

	if result.Translation != "" {
		return result.Translation, nil
	}

	return text, fmt.Errorf("no translation in response")
}

// DetectLanguage is a placeholder - language detection not currently used
func (c *Client) DetectLanguage(text string) (string, error) {
	return "en", nil
}
