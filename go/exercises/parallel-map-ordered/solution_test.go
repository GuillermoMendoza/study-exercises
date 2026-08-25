package parallelmapordered

import (
	"context"
	"reflect"
	"testing"
)

func TestParallelMapKeepsOrder(t *testing.T) {
	got, err := ParallelMap(context.Background(), []int{1, 2, 3}, 2, func(_ context.Context, n int) (int, error) { return n * 2, nil })
	if err != nil || !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Fatalf("got=%v err=%v", got, err)
	}
}
