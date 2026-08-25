package firstsuccessfulresult

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestReturnsFirstSuccess(t *testing.T) {
	got, err := FirstSuccessful(context.Background(), time.Second, []Task[string]{func(context.Context) (string, error) { return "", errors.New("no") }, func(context.Context) (string, error) { return "ok", nil }})
	if err != nil || got != "ok" {
		t.Fatalf("got=%q err=%v", got, err)
	}
}
