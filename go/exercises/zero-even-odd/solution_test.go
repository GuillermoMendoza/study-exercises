package zeroevenodd

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestZeroEvenOddPrintsInOrder(t *testing.T) {
	exercise := NewZeroEvenOdd(5)
	var output []int
	var outputMu sync.Mutex
	collect := func(value int) {
		outputMu.Lock()
		defer outputMu.Unlock()
		output = append(output, value)
	}

	var workers sync.WaitGroup
	workers.Add(3)
	go func() { defer workers.Done(); exercise.Zero(collect) }()
	go func() { defer workers.Done(); exercise.Even(collect) }()
	go func() { defer workers.Done(); exercise.Odd(collect) }()

	done := make(chan struct{})
	go func() { workers.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("workers did not finish")
	}

	want := []int{0, 1, 0, 2, 0, 3, 0, 4, 0, 5}
	if !reflect.DeepEqual(output, want) {
		t.Fatalf("output = %v, want %v", output, want)
	}
}
