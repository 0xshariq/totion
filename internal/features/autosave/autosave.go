package autosave

import (
	"time"
)

// AutoSaver manages automatic saving of notes
type AutoSaver struct {
	interval     time.Duration
	ticker       *time.Ticker
	stopChan     chan bool
	saveCallback func() error
}

// NewAutoSaver creates a new auto-saver with 30 second interval
func NewAutoSaver(saveCallback func() error) *AutoSaver {
	return &AutoSaver{
		interval:     30 * time.Second,
		stopChan:     make(chan bool),
		saveCallback: saveCallback,
	}
}

// Start begins the auto-save timer
func (a *AutoSaver) Start() {
	a.ticker = time.NewTicker(a.interval)
	go func() {
		for {
			select {
			case <-a.ticker.C:
				if a.saveCallback != nil {
					_ = a.saveCallback()
				}
			case <-a.stopChan:
				return
			}
		}
	}()
}

// Stop stops the auto-save timer
func (a *AutoSaver) Stop() {
	if a.ticker != nil {
		a.ticker.Stop()
	}
	a.stopChan <- true
}

// Reset resets the timer (call after manual save)
func (a *AutoSaver) Reset() {
	if a.ticker != nil {
		a.ticker.Reset(a.interval)
	}
}
