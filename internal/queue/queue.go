package queue

import "sync"

// Task is a minimal background task.
type Task struct {
	ID string
	Text string
	Tries int
}

// Queue is a tiny in-memory FIFO with mutex. Persistence is file-based in prod.
type Queue struct {
	mu sync.Mutex
	items []Task
}

func New() *Queue { return &Queue{} }

func (q *Queue) Add(t Task) { q.mu.Lock(); defer q.mu.Unlock(); q.items = append(q.items, t) }

func (q *Queue) Next() (Task, bool) {
	q.mu.Lock(); defer q.mu.Unlock()
	if len(q.items)==0 { return Task{}, false }
	t := q.items[0]; q.items = q.items[1:]; return t, true
}

func (q *Queue) Len() int { q.mu.Lock(); defer q.mu.Unlock(); return len(q.items) }
