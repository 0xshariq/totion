package linking

import (
	"regexp"
	"strings"
)

// Link represents a wiki-style link
type Link struct {
	Source string
	Target string
	Line   int
}

// LinkManager handles note linking and backlinks
type LinkManager struct {
	links map[string][]Link // Map of note -> outgoing links
}

// NewLinkManager creates a new link manager
func NewLinkManager() *LinkManager {
	return &LinkManager{
		links: make(map[string][]Link),
	}
}

// ParseLinks extracts wiki-style links from content
// Format: [[Note Title]] or [[Note Title|Display Text]]
func (lm *LinkManager) ParseLinks(content, sourceName string) []Link {
	links := []Link{}
	lines := strings.Split(content, "\n")

	// Regex to match [[link]] or [[link|text]]
	linkRegex := regexp.MustCompile(`\[\[([^\]|]+)(?:\|([^\]]+))?\]\]`)

	for i, line := range lines {
		matches := linkRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			target := strings.TrimSpace(match[1])
			links = append(links, Link{
				Source: sourceName,
				Target: target,
				Line:   i,
			})
		}
	}

	return links
}

// AddLinks adds links for a note
func (lm *LinkManager) AddLinks(noteName string, links []Link) {
	lm.links[noteName] = links
}

// GetOutgoingLinks returns outgoing links from a note
func (lm *LinkManager) GetOutgoingLinks(noteName string) []Link {
	return lm.links[noteName]
}

// GetBacklinks returns backlinks to a note
func (lm *LinkManager) GetBacklinks(noteName string) []Link {
	backlinks := []Link{}

	for _, links := range lm.links {
		for _, link := range links {
			if link.Target == noteName {
				backlinks = append(backlinks, link)
			}
		}
	}

	return backlinks
}

// CreateWikiLink creates a wiki-style link string
func (lm *LinkManager) CreateWikiLink(target string) string {
	return "[[" + target + "]]"
}

// CreateWikiLinkWithDisplay creates a wiki-style link with custom display text
func (lm *LinkManager) CreateWikiLinkWithDisplay(target, displayText string) string {
	return "[[" + target + "|" + displayText + "]]"
}

// IsWikiLink checks if text contains a wiki-style link
func (lm *LinkManager) IsWikiLink(text string) bool {
	linkRegex := regexp.MustCompile(`\[\[[^\]]+\]\]`)
	return linkRegex.MatchString(text)
}

// ExtractLinkTarget extracts the target from a wiki link
func (lm *LinkManager) ExtractLinkTarget(linkText string) string {
	linkRegex := regexp.MustCompile(`\[\[([^\]|]+)(?:\|[^\]]+)?\]\]`)
	matches := linkRegex.FindStringSubmatch(linkText)

	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}

	return ""
}
