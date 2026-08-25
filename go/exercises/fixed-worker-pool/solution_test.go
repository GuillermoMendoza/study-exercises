package fixedworkerpool

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestPoolCollectsOutcomesAndShutsDown(t *testing.T) {
	p := NewFixedWorkerPool[int](2)
	ok := p.Submit(func(context.Context) (int, error) { return 42, nil })
	bad := p.Submit(func(context.Context) (int, error) { return 0, errors.New("boom") })
	if r := <-ok; r.Value != 42 || r.Err != nil {
		t.Fatalf("success result = %#v", r)
	}
	if r := <-bad; r.Err == nil {
		t.Fatal("failed task lost its error")
	}
	p.Shutdown()
	if !p.AwaitTermination(time.Second) {
		t.Fatal("pool did not terminate")
	}
}
