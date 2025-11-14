package tasks

import (
	"regexp"
	"strings"
)

// Task represents a single task/checkbox
type Task struct {
	Text      string
	Completed bool
	Line      int
}

// TaskManager handles task operations
type TaskManager struct{}

// NewTaskManager creates a new task manager
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// ParseTasks extracts tasks from note content
func (tm *TaskManager) ParseTasks(content string) []Task {
	tasks := []Task{}
	lines := strings.Split(content, "\n")

	// Regex to match markdown checkboxes: - [ ] or - [x]
	checkboxRegex := regexp.MustCompile(`^[\s]*[-*]\s+\[([ xX])\]\s+(.+)$`)

	for i, line := range lines {
		if matches := checkboxRegex.FindStringSubmatch(line); matches != nil {
			completed := strings.ToLower(matches[1]) == "x"
			text := matches[2]

			tasks = append(tasks, Task{
				Text:      text,
				Completed: completed,
				Line:      i,
			})
		}
	}

	return tasks
}

// ToggleTask toggles a task's completion status
func (tm *TaskManager) ToggleTask(content string, lineNum int) string {
	lines := strings.Split(content, "\n")

	if lineNum < 0 || lineNum >= len(lines) {
		return content
	}

	line := lines[lineNum]

	// Toggle [ ] to [x] or [x] to [ ]
	if strings.Contains(line, "[ ]") {
		lines[lineNum] = strings.Replace(line, "[ ]", "[x]", 1)
	} else if strings.Contains(line, "[x]") || strings.Contains(line, "[X]") {
		lines[lineNum] = strings.Replace(strings.Replace(line, "[x]", "[ ]", 1), "[X]", "[ ]", 1)
	}

	return strings.Join(lines, "\n")
}

// GetTaskStats returns task statistics
func (tm *TaskManager) GetTaskStats(tasks []Task) (total, completed int) {
	total = len(tasks)
	for _, task := range tasks {
		if task.Completed {
			completed++
		}
	}
	return
}

// AddTask adds a new task to content
func (tm *TaskManager) AddTask(content, taskText string) string {
	newTask := "- [ ] " + taskText
	if content == "" {
		return newTask
	}
	return content + "\n" + newTask
}
