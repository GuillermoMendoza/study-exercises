package workerpool

import (
	"context"
	"errors"
	"sync"
)

type Task[T any] func(context.Context) (T, error)

type Result[T any] struct {
	Index int
	Value T
	Err   error
}

// Run executes at most workers tasks concurrently. Results preserve task order.
func Run[T any](ctx context.Context, workers int, tasks []Task[T]) ([]Result[T], error) {
	if workers <= 0 {
		return nil, errors.New("workerpool: workers must be greater than zero")
	}

	results := make([]Result[T], len(tasks))
	for i := range results {
		results[i].Index = i
	}
	if len(tasks) == 0 {
		return results, ctx.Err()
	}

	if workers > len(tasks) {
		workers = len(tasks)
	}

	type job struct {
		index int
		task  Task[T]
	}

	workCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	jobs := make(chan job) // Unbuffered: applies backpressure.

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			for j := range jobs {
				// Don't begin queued work after cancellation.
				if err := workCtx.Err(); err != nil {
					results[j.index].Err = err
					continue
				}

				if j.task == nil {
					results[j.index].Err = errors.New("workerpool: nil task")
					continue
				}

				results[j.index].Value, results[j.index].Err = j.task(workCtx)
			}
		}()
	}

	next := 0

producer:
	for ; next < len(tasks); next++ {
		if err := ctx.Err(); err != nil {
			break
		}

		select {
		case <-ctx.Done():
			break producer
		case jobs <- job{index: next, task: tasks[next]}:
			continue
		}
	}

	// Tasks never submitted are marked cancelled.
	if err := ctx.Err(); err != nil {
		for i := next; i < len(tasks); i++ {
			results[i].Err = err
		}
	}

	close(jobs)
	wg.Wait()

	return results, ctx.Err()
}
