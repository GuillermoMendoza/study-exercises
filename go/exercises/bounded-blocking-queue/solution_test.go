package boundedblockingqueue

import (
	"testing"
	"time"
)

func TestBoundedBlockingQueueKeepsFIFOAndUnblocksPut(t *testing.T) {
	queue := NewBoundedBlockingQueue[int](1)
	queue.Put(1)

	putFinished := make(chan struct{})
	go func() {
		queue.Put(2)
		close(putFinished)
	}()

	select {
	case <-putFinished:
		t.Fatal("Put returned while the queue was full")
	case <-time.After(100 * time.Millisecond):
	}

	if got := queue.Take(); got != 1 {
		t.Fatalf("first Take() = %d, want 1", got)
	}
	select {
	case <-putFinished:
	case <-time.After(time.Second):
		t.Fatal("Put did not unblock after Take")
	}
	if got := queue.Take(); got != 2 {
		t.Fatalf("second Take() = %d, want 2", got)
	}
}
