package readerswritercache

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestLoadsOnlyOnce(t *testing.T) {
	c := NewReadersWriterCache[string, int]()
	var loads atomic.Int32
	var w sync.WaitGroup
	for range 8 {
		w.Add(1)
		go func() {
			defer w.Done()
			if got := c.GetOrLoad("x", func(string) int { return int(loads.Add(1)) }); got != 1 {
				t.Errorf("got %d", got)
			}
		}()
	}
	w.Wait()
	if loads.Load() != 1 {
		t.Fatalf("loads=%d", loads.Load())
	}
}
