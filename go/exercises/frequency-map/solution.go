package frequencymap

import (
	"sync"
	"sync/atomic"
)

// Implement UnsafeFrequencyMap first to observe the race, then the other two safe variants.
type FrequencyMap interface {
	Increment(string)
	Count(string) int
	Snapshot() map[string]int
}

type UnsafeFrequencyMap struct {
	counts map[string]int
}

func (m *UnsafeFrequencyMap) Increment(key string) {
	if m.counts == nil {
		m.counts = make(map[string]int)
	}

	m.counts[key]++
}

func (m *UnsafeFrequencyMap) Count(key string) int {
	return m.counts[key]
}

func (m *UnsafeFrequencyMap) Snapshot() map[string]int {
	snapshot := make(map[string]int, len(m.counts))

	for key, count := range m.counts {
		snapshot[key] = count
	}

	return snapshot
}

type LockedFrequencyMap struct {
	mu     sync.RWMutex
	counts map[string]int
}

func (m *LockedFrequencyMap) Increment(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.counts == nil {
		m.counts = make(map[string]int)
	}

	m.counts[key]++
}

func (m *LockedFrequencyMap) Count(key string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.counts[key]
}

func (m *LockedFrequencyMap) Snapshot() map[string]int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	snapshot := make(map[string]int, len(m.counts))

	for key, count := range m.counts {
		snapshot[key] = count
	}

	return snapshot
}

type ConcurrentFrequencyMap struct {
	counts sync.Map
}

func (m *ConcurrentFrequencyMap) Increment(key string) {
	value, _ := m.counts.LoadOrStore(key, &atomic.Int64{})
	counter := value.(*atomic.Int64)

	counter.Add(1)
}

func (m *ConcurrentFrequencyMap) Count(key string) int {
	value, found := m.counts.Load(key)
	if !found {
		return 0
	}

	return int(value.(*atomic.Int64).Load())
}

func (m *ConcurrentFrequencyMap) Snapshot() map[string]int {
	snapshot := make(map[string]int)

	m.counts.Range(func(key, value any) bool {
		snapshot[key.(string)] = int(value.(*atomic.Int64).Load())
		return true
	})

	return snapshot
}
