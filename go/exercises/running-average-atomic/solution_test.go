package runningaverageatomic

import (
	"math"
	"sync"
	"testing"
)

func TestEmptyAverageIsZero(t *testing.T) {
	if got := NewAtomicRunningAverage().GetAverage(); got != 0 {
		t.Fatalf("GetAverage() = %v, want 0", got)
	}
}
func TestFinalAverageIsExactAfterConcurrentWritersFinish(t *testing.T) {
	average := NewAtomicRunningAverage()
	const writers = 12
	const additionsPerWorker = 2000
	ready := make(chan struct{}, writers)
	start := make(chan struct{})
	var workers sync.WaitGroup
	workers.Add(writers)
	for worker := range writers {
		value := float64(worker)
		go func() {
			defer workers.Done()
			ready <- struct{}{}
			<-start
			for range additionsPerWorker {
				average.AddNumber(value)
			}
		}()
	}
	for range writers {
		<-ready
	}
	close(start)
	workers.Wait()
	want := float64(writers-1) / 2
	if got := average.GetAverage(); math.Abs(got-want) > 1e-9 {
		t.Fatalf("GetAverage() = %v, want %v", got, want)
	}
}
