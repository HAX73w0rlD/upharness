package daemon

import (
	"fmt"
	"sync"
	"time"
)

// Daemon ist ein Hintergrund-Prozess, der die Queue verarbeitet.
type Daemon struct {
	mu      sync.Mutex
	Poll    time.Duration
	Stop    chan struct{}
	running bool
	tasks   []func() error
}

// New erstellt einen neuen Daemon.
func New() *Daemon {
	return &Daemon{Poll: 3 * time.Second, Stop: make(chan struct{})}
}

// Start lauft den Daemon im Hintergrund bis Stop geschlossen wird.
func (d *Daemon) Start() {
	d.mu.Lock();
	if d.running {
		d.mu.Unlock();
		return
	}
	d.running = true
	d.mu.Unlock();
	go d.loop()
}

// Stop beendet den Daemon sauber.
func (d *Daemon) Stop() {
	close(d.Stop)
	d.mu.Lock();
	d.running = false
	d.mu.Unlock();
}

// Running zuruckgeben, ob der Daemon aktiv ist.
func (d *Daemon) Running() bool {
	d.mu.Lock();
	defer d.mu.Unlock();
	return d.running
}

// AddTask fugt eine ausfuhrbare Aufgabe hinzu.
func (d *Daemon) AddTask(t func() error) {
	d.mu.Lock();
	d.tasks = append(d.tasks, t)
	d.mu.Unlock();
}

// loop ist die Hauptschleife des Daemons.
func (d *Daemon) loop() {
	ticker := time.NewTicker(d.Poll)
	defer ticker.Stop();
	for {
		select {
		case <-d.Stop:
			fmt.Println("Daemon gestoppt")
			return
		case <-ticker.C:
			d.runTasks()
		}
	}
}

// runTasks fuhrt alle pending Tasks aus.
func (d *Daemon) runTasks() {
	d.mu.Lock();
	tasks := d.tasks
	d.tasks = nil
	d.mu.Unlock();
	for _, t := range tasks {
		if err := t(); err != nil {
			fmt.Println("Task-Fehler:", err)
		}
	}
}
