package pomodoro

import (
	"time"
)

// PomodoroState represents the pomodoro timer state
type PomodoroState int

const (
	StateIdle PomodoroState = iota
	StateWork
	StateShortBreak
	StateLongBreak
)

// PomodoroTimer handles pomodoro time management
type PomodoroTimer struct {
	workDuration       time.Duration
	shortBreakDuration time.Duration
	longBreakDuration  time.Duration
	currentState       PomodoroState
	startTime          time.Time
	endTime            time.Time
	pomodorosCompleted int
}

// NewPomodoroTimer creates a new pomodoro timer
func NewPomodoroTimer() *PomodoroTimer {
	return &PomodoroTimer{
		workDuration:       25 * time.Minute,
		shortBreakDuration: 5 * time.Minute,
		longBreakDuration:  15 * time.Minute,
		currentState:       StateIdle,
		pomodorosCompleted: 0,
	}
}

// StartWork starts a work session
func (pt *PomodoroTimer) StartWork() {
	pt.currentState = StateWork
	pt.startTime = time.Now()
	pt.endTime = pt.startTime.Add(pt.workDuration)
}

// StartShortBreak starts a short break
func (pt *PomodoroTimer) StartShortBreak() {
	pt.currentState = StateShortBreak
	pt.startTime = time.Now()
	pt.endTime = pt.startTime.Add(pt.shortBreakDuration)
}

// StartLongBreak starts a long break
func (pt *PomodoroTimer) StartLongBreak() {
	pt.currentState = StateLongBreak
	pt.startTime = time.Now()
	pt.endTime = pt.startTime.Add(pt.longBreakDuration)
}

// Stop stops the timer
func (pt *PomodoroTimer) Stop() {
	if pt.currentState == StateWork {
		pt.pomodorosCompleted++
	}
	pt.currentState = StateIdle
}

// GetTimeRemaining returns the remaining time
func (pt *PomodoroTimer) GetTimeRemaining() time.Duration {
	if pt.currentState == StateIdle {
		return 0
	}

	remaining := time.Until(pt.endTime)
	if remaining < 0 {
		return 0
	}

	return remaining
}

// IsActive checks if timer is active
func (pt *PomodoroTimer) IsActive() bool {
	return pt.currentState != StateIdle
}

// GetState returns the current state
func (pt *PomodoroTimer) GetState() PomodoroState {
	return pt.currentState
}

// GetPomodorosCompleted returns completed pomodoros count
func (pt *PomodoroTimer) GetPomodorosCompleted() int {
	return pt.pomodorosCompleted
}

// FormatTimeRemaining formats remaining time as MM:SS
func (pt *PomodoroTimer) FormatTimeRemaining() string {
	remaining := pt.GetTimeRemaining()
	minutes := int(remaining.Minutes())
	seconds := int(remaining.Seconds()) % 60

	return formatDuration(minutes, seconds)
}

func formatDuration(minutes, seconds int) string {
	return padZero(minutes) + ":" + padZero(seconds)
}

func padZero(n int) string {
	if n < 10 {
		return "0" + string(rune(n+'0'))
	}
	return string(rune(n/10+'0')) + string(rune(n%10+'0'))
}
