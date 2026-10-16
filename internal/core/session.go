package core

import "time"

// Session represents a lightweight user session.
type Session struct {
	ID      string
	Model   string
	Created int64
	Active  bool
}

// NewSession creates a new session.
func NewSession(id, model string) *Session {
	return &Session{
		ID:      id,
		Model:   model,
		Created: time.Now().Unix(),
		Active:  true,
	}
}
.×ŸŠw