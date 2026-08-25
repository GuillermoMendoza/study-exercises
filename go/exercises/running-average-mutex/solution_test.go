package runningaveragemutex

import (
	"math"
	"sync"
	"sync/atomic"
	"testing"
)

func TestEmptyAverageIsZero(t *testing.T) {
	if got := NewMutexRunningAverage().GetAverage(); got != 0 {
		t.Fatalf("GetAverage() = %v, want 0", got)
	}
}

func TestAverageStaysCoherentDuringConcurrentUpdates(t *testing.T) {
	average := NewMutexRunningAverage()
	const writers = 12
	const additionsPerWriter = 2000
	ready := make(chan struct{}, writers+1)
	start := make(chan struct{})
	var writersDone sync.WaitGroup
	writersDone.Add(writers)
	var reading atomic.Bool
	reading.Store(true)
	var invalid atomic.Bool
	readerDone := make(chan struct{})
	go func() {
		defer close(readerDone)
		ready <- struct{}{}
		<-start
		for reading.Load() {
			snapshot := average.GetAverage()
			if math.IsNaN(snapshot) || math.IsInf(snapshot, 0) || snapshot < 0 || snapshot > writers-1 {
				invalid.Store(true)
			}
		}
	}()
	for worker := range writers {
		value := float64(worker)
		go func() {
			defer writersDone.Done()
			ready <- struct{}{}
			<-start
			for range additionsPerWriter {
				average.AddNumber(value)
			}
		}()
	}
	for range writers + 1 {
		<-ready
	}
	close(start)
	writersDone.Wait()
	reading.Store(false)
	<-readerDone
	if invalid.Load() {
		t.Fatal("observed an average outside the range of input values")
	}
	if want := float64(writers-1) / 2; math.Abs(average.GetAverage()-want) > 1e-9 {
		t.Fatalf("GetAverage() = %v, want %v", average.GetAverage(), want)
	}
}
