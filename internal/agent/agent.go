package agent

import "time"

// Status eines Agents.
type Status string

const (
	StatusIdle    Status = "idle"
	StatusRunning Status = "running"
	StatusDone    Status = "done"
	StatusFailed  Status = "failed"
)

// Task ist eine ausfuehrbare Einheit fuer einen Agenten.
type Task struct {
	ID        string
	Goal      string
	Status    Status
	Created   time.Time
	Updated   time.Time
	Result    string
	Error     string
}

// Agent fuehrt autonome Tasks aus.
type Agent struct {
	ID      string
	Name    string
	Status  Status
	Current *Task
	Queue   []Task
	Stop    chan struct{}
}

// NewAgent erstellt einen neuen Agenten.
func NewAgent(id, name string) *Agent {
	return &Agent{
		ID:    id,
		Name:  name,
		Status: StatusIdle,
		Queue:  []Task{},
		Stop:   make(chan struct{}),
	}
}

// AddTask fuegt einen Task zur Warteschlange hinzu.
func (a *Agent) AddTask(goal string) {
	a.Queue = append(a.Queue, Task{
		ID:      "task-" + time.Now().Format("150405"),
		Goal:    goal,
		Status:  StatusIdle,
		Created: time.Now(),
	})
}

// NextTask liefert den naechsten ausstehenden Task.
func (a *Agent) NextTask() (Task, bool) {
	if len(a.Queue) == 0 { return Task{}, false }
	t := a.Queue[0]; a.Queue = a.Queue[1:]; return t, true
}

// Run beginnt die Verarbeitung der Queue (blockierend).
func (a *Agent) Run() {
	for {
		select {
		case <-a.Stop:
			a.Status = StatusIdle
			a.Current = nil
			return
		default:
			t, ok := a.NextTask()
			if !ok {
				time.Sleep(500 * time.Millisecond)
				continue
			}
			a.Current = &t
			a.Status = StatusRunning
			// Task-Ausfuehrung wird hier durch Worker/Daemon ausgeloest.
			a.Status = StatusDone
			a.Current.Status = StatusDone
			a.Current = nil
		}
	}
}

// Stop beendet den Agenten sauber.
func (a *Agent) Stop() { close(a.Stop) }
