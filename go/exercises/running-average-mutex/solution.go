package runningaveragemutex

import "sync"

// MutexRunningAverage starts with the unsafe two-field implementation. Add a
// sync.Mutex and use it in both methods so sum and count are one snapshot.
type MutexRunningAverage struct {
	mu    sync.RWMutex
	sum   float64
	count int
}

func NewMutexRunningAverage() *MutexRunningAverage {
	return &MutexRunningAverage{}
}

func (r *MutexRunningAverage) AddNumber(number float64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.sum += number
	r.count++
}

func (r *MutexRunningAverage) GetAverage() float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.count == 0 {
		return 0.0
	}
	return r.sum / float64(r.count)
}
