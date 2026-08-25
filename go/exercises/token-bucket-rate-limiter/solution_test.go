package tokenbucketratelimiter

import (
	"testing"
	"time"
)

func TestBucketRefillsFromInjectedClock(t *testing.T) {
	now := time.Unix(0, 0)
	b := NewTokenBucket(2, 2, func() time.Time { return now })
	if !b.TryAcquire(2) {
		t.Fatal("initial capacity unavailable")
	}
	if b.TryAcquire(1) {
		t.Fatal("bucket over-issued")
	}
	now = now.Add(500 * time.Millisecond)
	if !b.TryAcquire(1) {
		t.Fatal("bucket did not refill")
	}
}
