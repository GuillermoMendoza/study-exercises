package safecounter

import "sync/atomic"

// SafeCounter must be safe for simultaneous use by multiple goroutines.
type SafeCounter struct {
	count atomic.Int64
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

func (c *SafeCounter) Increment() {
	c.count.Add(1)
}

func (c *SafeCounter) Value() int {
	return int(c.count.Load())
}
