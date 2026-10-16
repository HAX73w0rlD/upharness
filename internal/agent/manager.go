package agent

import "sync"

// Manager verwaltet alle laufenden Agenten.
type Manager struct {
	mu     sync.Mutex
	agents map[string]*Agent
}

// NewManager erstellt einen neuen Agent-Manager.
func NewManager() *Manager {
	return &Manager{agents: make(map[string]*Agent)}
}

// Register fuegt einen Agenten hinzu und startet ihn.
func (m *Manager) Register(a *Agent) {
	m.mu.Lock();
	defer m.mu.Unlock();
	m.agents[a.ID] = a;
	go a.Run();
}

// Get liefert einen Agenten zurueck, falls vorhanden.
func (m *Manager) Get(id string) (*Agent, bool) {
	m.mu.Lock();
	defer m.mu.Unlock();
	a, ok := m.agents[id];
	return a, ok;
}

// List gibt alle Agenten zurueck.
func (m *Manager) List() []*Agent {
	m.mu.Lock();
	defer m.mu.Unlock();
	out := make([]*Agent, 0, len(m.agents));
	for _, a := range m.agents { out = append(out, a) };
	return out;
}

// StopAll beendet alle Agenten.
func (m *Manager) StopAll() {
	m.mu.Lock();
	defer m.mu.Unlock();
	for _, a := range m.agents { a.Stop() };
	m.agents = make(map[string]*Agent);
}
