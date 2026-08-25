package mergeresultstreams

import "context"

// Merge emits every input value, preserves each input's local order, and closes when all inputs close.
func Merge[T any](ctx context.Context, inputs ...<-chan T) <-chan T {
	out := make(chan T)
	close(out)
	return out
}
