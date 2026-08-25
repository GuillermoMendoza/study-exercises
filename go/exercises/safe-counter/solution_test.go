package safecounter

import (
	"sync"
	"testing"
)

func TestSafeCounterPreservesConcurrentIncrements(t *testing.T) {
	counter := NewSafeCounter()
	const workers = 12
	const incrementsPerWorker = 1500

	ready := make(chan struct{}, workers)
	start := make(chan struct{})
	var workersDone sync.WaitGroup
	workersDone.Add(workers)
	for range workers {
		go func() {
			defer workersDone.Done()
			ready <- struct{}{}
			<-start
			for range incrementsPerWorker {
				counter.Increment()
			}
		}()
	}
	for range workers {
		<-ready
	}
	close(start)
	workersDone.Wait()

	if want := workers * incrementsPerWorker; counter.Value() != want {
		t.Fatalf("Value() = %d, want %d", counter.Value(), want)
	}
}
