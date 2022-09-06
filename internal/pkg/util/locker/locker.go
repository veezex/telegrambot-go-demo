package locker

import (
	"sync"
)

type Locker interface {
	RLock() func()
	Lock() func()
}

type locker struct {
	mu sync.RWMutex
}

func New() Locker {
	return &locker{
		mu: sync.RWMutex{},
	}
}

func (l *locker) RLock() func() {
	l.mu.RLock()

	return func() {
		l.mu.RUnlock()
	}
}

func (l *locker) Lock() func() {
	l.mu.Lock()

	return func() {
		l.mu.Unlock()
	}
}
