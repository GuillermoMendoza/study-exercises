package mergeresultstreams

import (
	"context"
	"testing"
)

func TestMergeEmitsEveryValue(t *testing.T) {
	a := make(chan int, 2)
	b := make(chan int, 1)
	a <- 1
	a <- 3
	close(a)
	b <- 2
	close(b)
	seen := map[int]bool{}
	for v := range Merge(context.Background(), a, b) {
		seen[v] = true
	}
	if len(seen) != 3 {
		t.Fatalf("seen=%v", seen)
	}
}
