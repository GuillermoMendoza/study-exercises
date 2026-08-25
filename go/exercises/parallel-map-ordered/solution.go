package parallelmapordered

import "context"

// ParallelMap limits mapper concurrency and returns values in input order.
func ParallelMap[T any, R any](ctx context.Context, input []T, max int, mapper func(context.Context, T) (R, error)) ([]R, error) {
	return nil, nil
}
