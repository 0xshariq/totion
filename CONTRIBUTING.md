# Contributing to Totion

Thank you for your interest in contributing to Totion! This document provides guidelines and instructions for contributing to the project.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [How to Contribute](#how-to-contribute)
- [Coding Standards](#coding-standards)
- [Testing Guidelines](#testing-guidelines)
- [Pull Request Process](#pull-request-process)
- [Feature Requests](#feature-requests)
- [Bug Reports](#bug-reports)
- [Community](#community)

## 🤝 Code of Conduct

### Our Pledge

We are committed to providing a welcoming and inspiring community for all. Please be respectful and constructive in your interactions.

### Expected Behavior

- Use welcoming and inclusive language
- Be respectful of differing viewpoints and experiences
- Gracefully accept constructive criticism
- Focus on what is best for the community
- Show empathy towards other community members

### Unacceptable Behavior

- Harassment, discriminatory comments, or personal attacks
- Trolling, insulting/derogatory comments
- Public or private harassment
- Publishing others' private information without permission
- Any conduct which could reasonably be considered inappropriate

## 🚀 Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.21 or higher** - [Download Go](https://golang.org/dl/)
- **Git** - [Download Git](https://git-scm.com/)
- **wkhtmltopdf** (optional, for PDF export) - [Download wkhtmltopdf](https://wkhtmltopdf.org/)

### Fork the Repository

1. Fork the repository on GitHub by clicking the "Fork" button
2. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/totion.git
   cd totion
   ```
3. Add the upstream repository:
   ```bash
   git remote add upstream https://github.com/0xshariq/totion.git
   ```

## 🛠️ Development Setup

### 1. Install Dependencies

```bash
# Download Go modules
go mod download

# Verify installation
go mod verify
```

### 2. Build the Project

```bash
# Build the application
make build

# Or manually
go build -o totion ./cmd/totion
```

### 3. Run the Application

```bash
# Run using makefile
make run

# Or run directly
./totion
```

### 4. Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### 5. Development Workflow

```bash
# Create a new branch for your feature
git checkout -b feature/your-feature-name

# Make your changes
# ...

# Build and test
make build
make run

# Commit your changes
git add .
git commit -m "feat: add your feature description"

# Push to your fork
git push origin feature/your-feature-name
```

## 📁 Project Structure

Understanding the project structure will help you contribute effectively:

```
totion/
├── cmd/
│   └── totion/              # Application entry point
│       └── main.go          # Main function
│
├── internal/
│   ├── app/                 # Main application logic
│   │   ├── app.go          # Core Model, Init, Update
│   │   ├── handlers.go     # Event handlers and business logic
│   │   └── views.go        # UI rendering functions
│   │
│   ├── models/             # Data models
│   │   └── note.go         # Note model and file formats
│   │
│   ├── storage/            # File system operations
│   │   └── storage.go      # CRUD operations for notes
│   │
│   ├── ui/                 # User interface components
│   │   ├── components/     # Reusable UI components
│   │   │   ├── editor.go
│   │   │   ├── list.go
│   │   │   └── input.go
│   │   ├── help/           # Help system
│   │   │   └── help.go
│   │   └── styles/         # Styling and colors
│   │       └── styles.go
│   │
│   ├── features/           # Advanced features (modular)
│   │   ├── autosave/       # Auto-save functionality
│   │   ├── export/         # Export to HTML/PDF/Text
│   │   ├── import/         # Import from Notion/Obsidian
│   │   ├── git/            # Git integration
│   │   ├── sync/           # Backup and sync
│   │   ├── stats/          # Statistics tracking
│   │   ├── linking/        # Wiki-style links
│   │   ├── templates/      # Note templates
│   │   └── recent/         # Recent notes tracking
│   │
│   ├── notebook/           # Folder/notebook management
│   │   └── notebook.go
│   │
│   └── themes/             # Theme system
│       └── themes.go
│
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── makefile                # Build automation
├── README.md               # Project documentation
├── CONTRIBUTING.md         # This file
└── .gitignore             # Git ignore rules
```

### Key Concepts

- **Bubble Tea Pattern**: The app uses the Elm Architecture (Model-View-Update)
  - `Model`: Application state
  - `View()`: Renders the UI
  - `Update()`: Handles messages and updates state

- **Feature Modules**: Each feature is self-contained in `internal/features/`
  - Easy to add new features
  - Clear separation of concerns
  - Minimal coupling with core app

## 🎯 How to Contribute

### Types of Contributions

We welcome various types of contributions:

1. **Bug Fixes** - Fix issues and improve stability
2. **New Features** - Add new functionality
3. **Documentation** - Improve or add documentation
4. **Code Quality** - Refactoring, optimization, cleanup
5. **Tests** - Add or improve test coverage
6. **UI/UX** - Improve user interface and experience

### Good First Issues

Look for issues labeled `good first issue` or `help wanted` on GitHub. These are great starting points for new contributors.

## 📝 Coding Standards

### Go Style Guide

Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) and [Effective Go](https://golang.org/doc/effective_go.html).

### Formatting

- Use `gofmt` to format your code:
  ```bash
  gofmt -w .
  ```

- Use `go vet` to check for common errors:
  ```bash
  go vet ./...
  ```

### Naming Conventions

- **Functions**: Use camelCase for unexported, PascalCase for exported
  ```go
  func handleKeyPress()  // unexported
  func NewExporter()     // exported
  ```

- **Variables**: Use descriptive names
  ```go
  // Good
  currentNote *models.Note
  recentManager *recent.RecentManager
  
  // Bad
  n *models.Note
  rm *recent.RecentManager
  ```

- **Constants**: Use PascalCase or ALL_CAPS
  ```go
  const ViewHome ViewState = iota
  const MAX_RECENT_NOTES = 10
  ```

### Code Organization

1. **Group imports** by standard library, external, and internal:
   ```go
   import (
       // Standard library
       "fmt"
       "os"
       
       // External packages
       tea "github.com/charmbracelet/bubbletea"
       
       // Internal packages
       "github.com/0xshariq/totion/internal/models"
   )
   ```

2. **Add comments** for exported functions:
   ```go
   // NewExporter creates a new exporter instance
   // It initializes the exporter with default settings
   func NewExporter() *Exporter {
       return &Exporter{}
   }
   ```

3. **Keep functions small** - Aim for functions under 50 lines
4. **Single responsibility** - Each function should do one thing well

### Error Handling

Always handle errors explicitly:

```go
// Good
file, err := os.Open(path)
if err != nil {
    return fmt.Errorf("failed to open file: %w", err)
}

// Bad
file, _ := os.Open(path)
```

## 🧪 Testing Guidelines

### Writing Tests

1. **Test file naming**: `*_test.go`
2. **Test function naming**: `TestFunctionName`
3. **Use table-driven tests** for multiple scenarios:

```go
func TestExporter_ExportToHTML(t *testing.T) {
    tests := []struct {
        name        string
        content     string
        title       string
        wantErr     bool
    }{
        {
            name:    "basic export",
            content: "# Hello\nWorld",
            title:   "Test Note",
            wantErr: false,
        },
        // More test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            e := NewExporter()
            err := e.ExportToHTML(tt.content, tt.title, "/tmp/test.html")
            if (err != nil) != tt.wantErr {
                t.Errorf("ExportToHTML() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./internal/features/export

# Run with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🔄 Pull Request Process

### Before Submitting

1. **Update your branch** with latest upstream:
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```

2. **Run tests** and ensure they pass:
   ```bash
   go test ./...
   ```

3. **Format your code**:
   ```bash
   gofmt -w .
   go vet ./...
   ```

4. **Test the application** manually:
   ```bash
   make build
   make run
   ```

### Commit Messages

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

**Examples:**
```
feat(export): add PDF export functionality

Implemented PDF export using wkhtmltopdf library.
Users can now export notes to professional PDF format.

Closes #123
```

```
fix(editor): resolve auto-save timing issue

Fixed race condition in auto-save timer that caused
occasional data loss. Auto-save now properly resets
after manual saves.

Fixes #456
```

### PR Title and Description

**Title:** Use the same format as commit messages

**Description should include:**
- What changes were made
- Why the changes were made
- How to test the changes
- Screenshots (if UI changes)
- Related issues

**Template:**
```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Code refactoring

## Changes Made
- Change 1
- Change 2
- Change 3

## Testing
How to test these changes:
1. Step 1
2. Step 2
3. Expected result

## Screenshots (if applicable)
[Add screenshots here]

## Related Issues
Closes #123
Related to #456

## Checklist
- [ ] Code follows project style guidelines
- [ ] Tests added/updated and passing
- [ ] Documentation updated
- [ ] No breaking changes (or documented)
```

### Review Process

1. **Submit PR** on GitHub
2. **CI checks** will run automatically
3. **Maintainers review** - may request changes
4. **Address feedback** - make requested changes
5. **Approval** - maintainer approves PR
6. **Merge** - maintainer merges to main branch

## 💡 Feature Requests

### Proposing New Features

1. **Check existing issues** - ensure feature hasn't been requested
2. **Open a new issue** with label `enhancement`
3. **Describe the feature**:
   - What problem does it solve?
   - Who would benefit from it?
   - How should it work?
   - Are there any alternatives?

### Feature Template

```markdown
## Feature Description
Clear description of the proposed feature

## Problem Statement
What problem does this feature solve?

## Proposed Solution
How should this feature work?

## Alternatives Considered
What other solutions did you consider?

## Additional Context
Any other relevant information, screenshots, mockups, etc.
```

### Priority Features to Implement

Check the [README.md](README.md) roadmap section for high-priority features:
- Auto-save functionality ⭐
- Full-text search ⭐
- Tag system ⭐
- Task checkboxes ⭐
- Recently opened notes ⭐

## 🐛 Bug Reports

### Reporting Bugs

1. **Check existing issues** - bug might already be reported
2. **Open a new issue** with label `bug`
3. **Include details**:
   - Steps to reproduce
   - Expected behavior
   - Actual behavior
   - System information (OS, Go version)
   - Error messages or logs

### Bug Report Template

```markdown
## Bug Description
Clear description of the bug

## Steps to Reproduce
1. Step 1
2. Step 2
3. Step 3

## Expected Behavior
What should happen

## Actual Behavior
What actually happens

## System Information
- OS: [e.g., Ubuntu 22.04]
- Go Version: [e.g., 1.21.3]
- Totion Version: [e.g., v0.1.0]

## Error Messages
```
Paste error messages or logs here
```

## Screenshots
[Add screenshots if applicable]
```

## 🌟 Adding New Features

### Feature Module Pattern

When adding a new feature, follow this structure:

1. **Create feature directory**:
   ```
   internal/features/your-feature/
   └── your_feature.go
   ```

2. **Define feature interface**:
   ```go
   package yourfeature
   
   type YourFeatureManager struct {
       // fields
   }
   
   func NewYourFeatureManager() *YourFeatureManager {
       return &YourFeatureManager{}
   }
   
   func (m *YourFeatureManager) DoSomething() error {
       // implementation
       return nil
   }
   ```

3. **Integrate with app**:
   - Add to `Model` struct in `internal/app/app.go`
   - Add keyboard shortcut in `handlers.go`
   - Add view rendering in `views.go`

4. **Update documentation**:
   - Add to README.md features section
   - Update keyboard shortcuts table
   - Add help menu entry

### Example: Adding a New Feature

Let's say you want to add a "Favorites" feature:

```go
// internal/features/favorites/favorites.go
package favorites

import "github.com/0xshariq/totion/internal/models"

type FavoritesManager struct {
    favorites []string
}

func NewFavoritesManager() *FavoritesManager {
    return &FavoritesManager{
        favorites: []string{},
    }
}

func (f *FavoritesManager) AddFavorite(notePath string) {
    f.favorites = append(f.favorites, notePath)
}

func (f *FavoritesManager) GetFavorites() []string {
    return f.favorites
}
```

Then integrate in `app.go`, `handlers.go`, and `views.go`.

## 📚 Documentation

### README Updates

When adding features, update:
- Features list
- Keyboard shortcuts table
- Quick start guide (if applicable)
- Troubleshooting section (if applicable)

### Code Comments

- Add comments for exported functions
- Explain complex logic
- Use TODO comments for future improvements:
  ```go
  // TODO: Implement caching for better performance
  ```

### Help System

Update `internal/ui/help/help.go` with new features:
- Add to appropriate help topic
- Include keyboard shortcuts
- Provide usage examples

## 🤝 Community

### Getting Help

- **GitHub Issues** - Ask questions, report bugs
- **Discussions** - General discussions, ideas, Q&A
- **Pull Requests** - Code review and collaboration

### Communication

- Be respectful and professional
- Provide context and details
- Be patient - maintainers are volunteers
- Help others when you can

## 🎓 Learning Resources

### Bubble Tea Framework

- [Bubble Tea Tutorial](https://github.com/charmbracelet/bubbletea/tree/master/tutorials)
- [Bubble Tea Examples](https://github.com/charmbracelet/bubbletea/tree/master/examples)
- [Lipgloss Styling](https://github.com/charmbracelet/lipgloss)

### Go Programming

- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Tour](https://tour.golang.org/)

## 📄 License

By contributing to Totion, you agree that your contributions will be licensed under the MIT License.

---

## 🙏 Thank You!

Thank you for contributing to Totion! Every contribution, no matter how small, helps make this project better.

**Questions?** Feel free to open an issue or start a discussion on GitHub.

**Happy Coding! 🚀**
