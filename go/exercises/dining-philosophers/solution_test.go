package diningphilosophers

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestEveryoneEats(t *testing.T) {
	d := NewDiningPhilosophers()
	var ate atomic.Int32
	var w sync.WaitGroup
	for i := range 5 {
		w.Add(1)
		go func(id int) {
			defer w.Done()
			d.WantsToEat(id, func() {}, func() {}, func() { ate.Add(1) }, func() {}, func() {})
		}(i)
	}
	done := make(chan struct{})
	go func() { w.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("deadlock")
	}
	if ate.Load() != 5 {
		t.Fatalf("ate=%d", ate.Load())
	}
}
