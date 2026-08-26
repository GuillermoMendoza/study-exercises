package blockingqueuechannels

import "sync"

type BlockingQueue[T any] struct {
	items chan T
	once  sync.Once
}

func NewBlockingQueue[T any](capacity int) *BlockingQueue[T] {
	if capacity < 0 {
		panic("Capacity must be larger than 0")
	}

	return &BlockingQueue[T]{
		items: make(chan T, capacity),
	}
}

func (q *BlockingQueue[T]) Put(value T) {
	q.items <- value
}

func (q *BlockingQueue[T]) Get() (T, bool) {
	value, ok := <-q.items
	return value, ok
}

func (q *BlockingQueue[T]) Size() int {
	return len(q.items)
}

func (q *BlockingQueue[T]) Close() {
	q.once.Do(func() {
		close(q.items)
	})
}
