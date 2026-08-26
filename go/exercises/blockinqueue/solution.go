package blockinqueue

import "sync"

type BlockingQueue[T any] struct {
	mu       sync.RWMutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
	capacity int
	items    []T
	closed   bool
}

func NewBlockingQueue[T any](capacity int) *BlockingQueue[T] {
	if capacity < 0 {
		panic("Capacity must be larger than 0")
	}

	q := &BlockingQueue[T]{
		capacity: capacity,
		items:    make([]T, capacity),
	}

	q.notEmpty = sync.NewCond(&q.mu)
	q.notFull = sync.NewCond(&q.mu)

	return q
}

func (q *BlockingQueue[T]) Put(value T) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.items) == q.capacity && !q.closed {
		q.notFull.Wait()
	}

	if q.closed {
		return false
	}

	q.items = append(q.items, value)

	q.notEmpty.Signal()
	return true
}

func (q *BlockingQueue[T]) Get() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.items) == 0 && !q.closed {
		q.notEmpty.Wait()
	}

	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	value := q.items[0]

	var zero T
	q.items[0] = zero
	q.items = q.items[1:]

	q.notFull.Signal()

	return value, true
}

func (q *BlockingQueue[T]) Size() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return len(q.items)
}

func (q *BlockingQueue[T]) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return
	}

	q.closed = true

	q.notEmpty.Broadcast()
	q.notFull.Broadcast()
}
