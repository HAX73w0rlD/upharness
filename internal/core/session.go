package core

import (
	"sync"
	"time"
)

// Session represents a lightweight user session with model assignment.
type Session struct {
	ID      string
	Model   string
	Created int64
	Active  bool
	Updated int64
}

// Manager manages all user sessions.
type Manager struct {
	mu      sync.Mutex
	sessions map[string]*Session
}

// NewManager creates a new session manager.
func NewManager() *Manager {
	return &Manager{sessions: make(map[string]*Session)}
}

// Create creates or returns an existing session.
func (m *Manager) Create(id, model string) *Session {
	m.mu.Lock();
	defer m.mu.Unlock();
	if s, ok := m.sessions[id]; ok {
		s.Model = model
		s.Active = true
		s.Updated = time.Now().Unix()
		return s
	}
	s := &Session{ID: id, Model: model, Created: time.Now().Unix(), Active: true, Updated: time.Now().Unix()}
	m.sessions[id] = s
	return s
}

// Get returns a session by ID, or nil if not found.
func (m *Manager) Get(id string) *Session {
	m.mu.Lock();
	defer m.mu.Unlock();
	return m.sessions[id]
}

// List returns all active sessions.
func (m *Manager) List() []*Session {
	m.mu.Lock();
	defer m.mu.Unlock();
	out := make([]*Session, 0, len(m.sessions));
	for _, s := range m.sessions {
		if s.Active { out = append(out, s) }
	}
	return out;
}

// Close marks a session as inactive.
func (m *Manager) Close(id string) bool {
	m.mu.Lock();
	defer m.mu.Unlock();
	if s, ok := m.sessions[id]; ok {
		s.Active = false
		s.Updated = time.Now().Unix()
		return true
	}
	return false;
}

// SetModel changes the model of an active session.
func (m *Manager) SetModel(id, model string) bool {
	m.mu.Lock();
	defer m.mu.Unlock();
	if s, ok := m.sessions[id]; ok && s.Active {
		s.Model = model
		s.Updated = time.Now().Unix()
		return true
	}
	return false;
}
