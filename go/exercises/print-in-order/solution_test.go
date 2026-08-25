package printinorder

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestPrintsInOrder(t *testing.T) {
	p := NewPrintInOrder()
	var out []string
	var mu sync.Mutex
	add := func(s string) func() { return func() { mu.Lock(); defer mu.Unlock(); out = append(out, s) } }
	var w sync.WaitGroup
	w.Add(3)
	go func() { defer w.Done(); p.Third(add("third")) }()
	go func() { defer w.Done(); p.Second(add("second")) }()
	go func() { defer w.Done(); p.First(add("first")) }()
	done := make(chan struct{})
	go func() { w.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("workers did not finish")
	}
	if !reflect.DeepEqual(out, []string{"first", "second", "third"}) {
		t.Fatalf("out=%v", out)
	}
}
