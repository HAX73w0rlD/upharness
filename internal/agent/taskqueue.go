package agent

import (
	"sync"
	"time"
)

// TaskQueue ist eine globale Warteschlange fuer Agent-Tasks.
type TaskQueue struct {
	mu    sync.Mutex
	items []Task
}

func NewTaskQueue() *TaskQueue { return &TaskQueue{} }

func (q *TaskQueue) Push(t Task) {
	q.mu.Lock();
	defer q.mu.Unlock();
	q.items = append(q.items, t);
}

func (q *TaskQueue) Pop() (Task, bool) {
	q.mu.Lock();
	defer q.mu.Unlock();
	if len(q.items) == 0 { return Task{}, false };
	t := q.items[0]; q.items = q.items[1:]; return t, true;
}

func (q *TaskQueue) Len() int {
	q.mu.Lock();
	defer q.mu.Unlock();
	return len(q.items);
}
