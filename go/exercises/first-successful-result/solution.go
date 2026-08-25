package firstsuccessfulresult

import (
	"context"
	"errors"
	"time"
)

type Task[T any] func(context.Context) (T, error)

func FirstSuccessful[T any](ctx context.Context, timeout time.Duration, tasks []Task[T]) (T, error) {
	var zero T

	if len(tasks) == 0 {
		return zero, errors.New("No tasks provided")
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	resChan := make(chan T, 1)
	errChan := make(chan error, len(tasks))

	for _, task := range tasks {
		go func(parallelTask Task[T]) {
			res, err := parallelTask(ctx)

			if err != nil {
				errChan <- err
				return
			}

			select {
			case resChan <- res:
			default:
			}
		}(task)
	}

	failedCount := 0

	for {
		select {
		case res := <-resChan:
			return res, nil
		case <-errChan:
			failedCount++
			if failedCount == len(tasks) {
				return zero, errors.New("All tasks failed")
			}
		case <-ctx.Done():
			return zero, ctx.Err()
		}
	}

}
