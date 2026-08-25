package safecountermutex

import "sync"

// MutexSafeCounter must use sync.Mutex to protect its shared counter state.
// Do not use sync/atomic for this exercise.
type MutexSafeCounter struct {
	mu    sync.Mutex
	count int
}

func NewMutexSafeCounter() *MutexSafeCounter {
	return &MutexSafeCounter{}
}

func (c *MutexSafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *MutexSafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
