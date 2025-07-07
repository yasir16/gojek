package schema

import (
	"time"
)

const (
	// DefaultPrompt is string prefix.
	DefaultPrompt = ""
)

// IShellHistory holds all the commands history
type IShellHistory struct {
	Command   string
	CreatedAt time.Time
}

// IShell is an interactive cli shell.
type IShell struct {
	History []*IShellHistory
}

// RecordHistory holds the History for shell session
func (s *IShell) RecordHistory(cmd string) {
	record := &IShellHistory{
		Command:   cmd,
		CreatedAt: time.Now(),
	}
	s.History = append(s.History, record)
}

// RecordShellHistory holds the History for shell session
func (c *Command) RecordShellHistory(history []*IShellHistory) {
	c.ShellHistory = history
}

func (s *IShell) ShowPrompt() {

}
