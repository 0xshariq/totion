package lingo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client handles communication with Lingo.dev API directly
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// Response represents the API response structure
type Response struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   string          `json:"error,omitempty"`
}

// NewClient creates a new Lingo.dev client that calls the API directly
// No need for Node.js bridge server!
func NewClient(apiKey string) *Client {
	if apiKey == "" {
		// Return a disabled client if no API key
		return &Client{
			apiKey:  "",
			baseURL: "https://api.lingo.dev/v1", // Official API endpoint
			httpClient: &http.Client{
				Timeout: 30 * time.Second,
			},
		}
	}

	return &Client{
		apiKey:  apiKey,
		baseURL: "https://api.lingo.dev/v1",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// IsEnabled returns true if the client has an API key configured
func (c *Client) IsEnabled() bool {
	return c.apiKey != ""
}

// TranslateText translates a text string to target language
// Calls Lingo.dev API directly - no bridge server needed!
func (c *Client) TranslateText(text, sourceLocale, targetLocale string, fast bool) (string, error) {
	if !c.IsEnabled() {
		return "", fmt.Errorf("translation disabled: no API key configured")
	}

	payload := map[string]interface{}{
		"text":         text,
		"sourceLocale": sourceLocale,
		"targetLocale": targetLocale,
	}
	if fast {
		payload["fast"] = true
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/translate", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// Extract translation from response
	if translation, ok := result["translation"].(string); ok {
		return translation, nil
	}

	return "", fmt.Errorf("unexpected response format")
}

// TranslateBatch translates text to multiple languages at once
// Calls Lingo.dev API directly - no bridge server needed!
func (c *Client) TranslateBatch(text, sourceLocale string, targetLocales []string) ([]string, error) {
	if !c.IsEnabled() {
		return nil, fmt.Errorf("translation disabled: no API key configured")
	}

	payload := map[string]interface{}{
		"text":          text,
		"sourceLocale":  sourceLocale,
		"targetLocales": targetLocales,
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
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if translations, ok := result["translations"].([]interface{}); ok {
		result := make([]string, len(translations))
		for i, t := range translations {
			if str, ok := t.(string); ok {
				result[i] = str
			}
		}
		return result, nil
	}

	return nil, fmt.Errorf("unexpected response format")
}

// TranslateHTML translates HTML/Markdown while preserving formatting
// Calls Lingo.dev API directly - no bridge server needed!
func (c *Client) TranslateHTML(html, sourceLocale, targetLocale string) (string, error) {
	if !c.IsEnabled() {
		return "", fmt.Errorf("translation disabled: no API key configured")
	}

	payload := map[string]interface{}{
		"html":         html,
		"sourceLocale": sourceLocale,
		"targetLocale": targetLocale,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/translate/html", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if translation, ok := result["translation"].(string); ok {
		return translation, nil
	}

	return "", fmt.Errorf("unexpected response format")
}

// DetectLanguage detects the language of the given text
// Calls Lingo.dev API directly - no bridge server needed!
func (c *Client) DetectLanguage(text string) (string, error) {
	if !c.IsEnabled() {
		return "en", nil // Default to English if no API key
	}

	payload := map[string]interface{}{
		"text": text,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/detect", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if locale, ok := result["locale"].(string); ok {
		return locale, nil
	}

	return "en", nil // Default fallback
}

// TranslateObject translates an entire object (for note metadata)
// Note: This may not be available in all Lingo.dev API versions
func (c *Client) TranslateObject(obj map[string]interface{}, sourceLocale, targetLocale string) (map[string]interface{}, error) {
	if !c.IsEnabled() {
		return nil, fmt.Errorf("translation disabled: no API key configured")
	}

	// Recursively translate all string values in the object
	result := make(map[string]interface{})
	for key, value := range obj {
		if str, ok := value.(string); ok {
			translated, err := c.TranslateText(str, sourceLocale, targetLocale, false)
			if err != nil {
				result[key] = value // Keep original on error
			} else {
				result[key] = translated
			}
		} else {
			result[key] = value
		}
	}

	return result, nil
}
