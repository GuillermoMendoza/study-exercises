package fixedworkerpool

import (
	"context"
	"time"
)

type Task[T any] func(context.Context) (T, error)

type TaskResult[T any] struct {
	Value T
	Err   error
}

// FixedWorkerPool must run no more than its configured number of tasks at once.
type FixedWorkerPool[T any] struct{}

func NewFixedWorkerPool[T any](workers int) *FixedWorkerPool[T] {
	if workers < 1 {
		panic("workers must be positive")
	}
	return &FixedWorkerPool[T]{}
}
func (p *FixedWorkerPool[T]) Submit(task Task[T]) <-chan TaskResult[T] {
	result := make(chan TaskResult[T], 1)
	result <- TaskResult[T]{}
	close(result)
	return result
}
func (p *FixedWorkerPool[T]) Shutdown()                                   {}
func (p *FixedWorkerPool[T]) AwaitTermination(timeout time.Duration) bool { return false }
