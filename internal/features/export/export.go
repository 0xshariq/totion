package export

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// ExportFormat represents the export format type
type ExportFormat string

const (
	FormatHTML      ExportFormat = "html"
	FormatPDF       ExportFormat = "pdf"
	FormatPlainText ExportFormat = "txt"
	FormatMarkdown  ExportFormat = "md"
	FormatJSON      ExportFormat = "json"
)

// Exporter handles exporting notes to various formats
type Exporter struct{}

// NewExporter creates a new exporter
func NewExporter() *Exporter {
	return &Exporter{}
}

// ExportToHTML exports markdown content to HTML
func (e *Exporter) ExportToHTML(content, title, outputPath string) error {
	htmlTemplate := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif;
            line-height: 1.6;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            color: #333;
        }
        h1, h2, h3 { color: #2c3e50; }
        code {
            background: #f4f4f4;
            padding: 2px 5px;
            border-radius: 3px;
        }
        pre {
            background: #f4f4f4;
            padding: 15px;
            border-radius: 5px;
            overflow-x: auto;
        }
    </style>
</head>
<body>
    <h1>{{.Title}}</h1>
    <pre>{{.Content}}</pre>
</body>
</html>`

	tmpl, err := template.New("export").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	data := struct {
		Title   string
		Content string
	}{
		Title:   title,
		Content: content,
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}

// ExportToPlainText exports content to plain text
func (e *Exporter) ExportToPlainText(content, outputPath string) error {
	// Remove markdown formatting for plain text
	plainContent := e.stripMarkdown(content)

	if err := os.WriteFile(outputPath, []byte(plainContent), 0644); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

// ExportToMarkdown copies markdown content
func (e *Exporter) ExportToMarkdown(content, outputPath string) error {
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

// ExportToPDF exports content to PDF format using wkhtmltopdf
func (e *Exporter) ExportToPDF(content, title, outputPath string) error {
	// Check if wkhtmltopdf is installed
	if _, err := exec.LookPath("wkhtmltopdf"); err != nil {
		return fmt.Errorf("wkhtmltopdf not installed. Install it with: sudo apt-get install wkhtmltopdf")
	}

	// Generate HTML first with proper styling for PDF
	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>%s</title>
    <style>
        @page { size: A4; margin: 2cm; }
        body {
            font-family: 'Georgia', 'Times New Roman', serif;
            line-height: 1.8;
            color: #000;
            font-size: 12pt;
            max-width: 800px;
            margin: 0 auto;
        }
        h1 {
            color: #1a1a1a;
            border-bottom: 2px solid #333;
            padding-bottom: 10px;
            margin-bottom: 20px;
            font-size: 24pt;
        }
        h2 { 
            color: #2a2a2a; 
            margin-top: 20px;
            font-size: 18pt;
        }
        h3 { 
            color: #3a3a3a;
            font-size: 14pt;
        }
        code {
            background: #f0f0f0;
            padding: 2px 6px;
            border-radius: 3px;
            font-family: 'Courier New', monospace;
            font-size: 10pt;
        }
        pre {
            background: #f5f5f5;
            padding: 15px;
            border-left: 4px solid #333;
            overflow-x: auto;
            font-family: 'Courier New', monospace;
            font-size: 9pt;
            line-height: 1.4;
        }
        pre code {
            background: transparent;
            padding: 0;
        }
        blockquote {
            border-left: 4px solid #ddd;
            padding-left: 15px;
            color: #666;
            margin: 15px 0;
            font-style: italic;
        }
        ul, ol { 
            margin: 10px 0; 
            padding-left: 30px; 
        }
        li { 
            margin: 5px 0; 
        }
        a { 
            color: #0066cc; 
            text-decoration: underline; 
        }
        table {
            border-collapse: collapse;
            width: 100%%;
            margin: 15px 0;
        }
        th, td {
            border: 1px solid #ddd;
            padding: 8px;
            text-align: left;
        }
        th { 
            background-color: #f0f0f0; 
            font-weight: bold; 
        }
        .footer {
            margin-top: 40px;
            padding-top: 20px;
            border-top: 1px solid #ddd;
            font-size: 9pt;
            color: #666;
            text-align: center;
        }
    </style>
</head>
<body>
    <h1>%s</h1>
    <div>%s</div>
    <div class="footer">
        Generated by Totion
    </div>
</body>
</html>`, title, title, e.convertMarkdownToHTML(content))

	// Create temporary HTML file
	tmpHTMLPath := strings.TrimSuffix(outputPath, ".pdf") + "_temp.html"
	if err := os.WriteFile(tmpHTMLPath, []byte(htmlContent), 0644); err != nil {
		return fmt.Errorf("error writing temporary HTML file: %w", err)
	}
	defer os.Remove(tmpHTMLPath) // Clean up temp file

	// Convert HTML to PDF using wkhtmltopdf
	cmd := exec.Command("wkhtmltopdf",
		"--encoding", "UTF-8",
		"--page-size", "A4",
		"--margin-top", "20mm",
		"--margin-bottom", "20mm",
		"--margin-left", "20mm",
		"--margin-right", "20mm",
		"--enable-local-file-access",
		tmpHTMLPath,
		outputPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error converting to PDF: %w\nOutput: %s", err, string(output))
	}

	return nil
}

// convertMarkdownToHTML converts basic markdown to HTML
func (e *Exporter) convertMarkdownToHTML(content string) string {
	html := content

	// Headers
	for i := 6; i >= 1; i-- {
		prefix := strings.Repeat("#", i)
		html = strings.ReplaceAll(html, prefix+" ", fmt.Sprintf("<h%d>", i))
		html = strings.ReplaceAll(html, "\n", fmt.Sprintf("</h%d>\n", i))
	}

	// Bold
	html = strings.ReplaceAll(html, "**", "<strong>")
	html = strings.ReplaceAll(html, "**", "</strong>")

	// Italic
	html = strings.ReplaceAll(html, "*", "<em>")
	html = strings.ReplaceAll(html, "*", "</em>")

	// Code
	html = strings.ReplaceAll(html, "`", "<code>")
	html = strings.ReplaceAll(html, "`", "</code>")

	// Line breaks
	html = strings.ReplaceAll(html, "\n\n", "</p><p>")
	html = "<p>" + html + "</p>"

	return html
}

// stripMarkdown removes basic markdown formatting
func (e *Exporter) stripMarkdown(content string) string {
	// Simple markdown stripping (can be enhanced)
	result := content

	// Remove headers
	result = strings.ReplaceAll(result, "#", "")

	// Remove bold/italic
	result = strings.ReplaceAll(result, "**", "")
	result = strings.ReplaceAll(result, "*", "")
	result = strings.ReplaceAll(result, "__", "")
	result = strings.ReplaceAll(result, "_", "")

	// Remove inline code
	result = strings.ReplaceAll(result, "`", "")

	return result
}

// GetExportFormats returns available export formats
func (e *Exporter) GetExportFormats() []ExportFormat {
	return []ExportFormat{
		FormatHTML,
		FormatPDF,
		FormatPlainText,
		FormatMarkdown,
		FormatJSON,
	}
}

// ExportToJSON exports note metadata and content as JSON
func (e *Exporter) ExportToJSON(content, title, outputPath string) error {
	data := map[string]interface{}{
		"title":       title,
		"content":     content,
		"exported_at": time.Now().Format(time.RFC3339),
		"format":      "markdown",
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	if err := os.WriteFile(outputPath, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

// BatchExport exports multiple notes to a directory
func (e *Exporter) BatchExport(notes []NoteData, outputDir string, format ExportFormat) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	for _, note := range notes {
		filename := note.Title
		var err error

		switch format {
		case FormatHTML:
			outputPath := filepath.Join(outputDir, filename+".html")
			err = e.ExportToHTML(note.Content, note.Title, outputPath)
		case FormatPDF:
			outputPath := filepath.Join(outputDir, filename+".pdf")
			err = e.ExportToPDF(note.Content, note.Title, outputPath)
		case FormatPlainText:
			outputPath := filepath.Join(outputDir, filename+".txt")
			err = e.ExportToPlainText(note.Content, outputPath)
		case FormatMarkdown:
			outputPath := filepath.Join(outputDir, filename+".md")
			err = e.ExportToMarkdown(note.Content, outputPath)
		case FormatJSON:
			outputPath := filepath.Join(outputDir, filename+".json")
			err = e.ExportToJSON(note.Content, note.Title, outputPath)
		}

		if err != nil {
			return fmt.Errorf("error exporting %s: %w", note.Title, err)
		}
	}

	return nil
}

// NoteData represents note data for batch export
type NoteData struct {
	Title   string
	Content string
	Path    string
}
