package boundedblockingqueue

// BoundedBlockingQueue is a fixed-capacity FIFO queue. Put and Take must block
// until they can make progress.
type BoundedBlockingQueue[T any] struct {
	ch chan T
}

func NewBoundedBlockingQueue[T any](capacity int) *BoundedBlockingQueue[T] {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	// TODO: retain capacity and initialize queue state.
	return &BoundedBlockingQueue[T]{
		ch: make(chan T, capacity),
	}
}

func (q *BoundedBlockingQueue[T]) Put(value T) {
	q.ch <- value
}

func (q *BoundedBlockingQueue[T]) Take() T {
	// TODO: wait while empty, then remove and return the oldest value.
	return <-q.ch
}
