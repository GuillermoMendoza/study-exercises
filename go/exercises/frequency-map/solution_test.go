package frequencymap

import (
	"sync"
	"testing"
)

func TestLockedFrequencyMap(t *testing.T) { assertCounts(t, &LockedFrequencyMap{}) }

func TestConcurrentFrequencyMap(t *testing.T) { assertCounts(t, &ConcurrentFrequencyMap{}) }

func assertCounts(t *testing.T, m FrequencyMap) {
	t.Helper()
	var w sync.WaitGroup
	for range 8 {
		w.Add(1)
		go func() {
			defer w.Done()
			for range 1000 {
				m.Increment("go")
			}
		}()
	}
	w.Wait()
	if got := m.Count("go"); got != 8000 {
		t.Fatalf("Count() = %d, want 8000", got)
	}
}
