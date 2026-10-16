package daemon

import "time"

// Daemon is a minimal poll loop stub.
type Daemon struct {
	Poll time.Duration
	Stop chan struct{}
}

func New() *Daemon { return &Daemon{Poll: 3 * time.Second, Stop: make(chan struct{})} }
