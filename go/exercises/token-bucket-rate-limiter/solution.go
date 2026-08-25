package tokenbucketratelimiter

import "time"

// TokenBucket must guard refill and acquisition with one mutex.
type TokenBucket struct{}

func NewTokenBucket(capacity int, rate float64, now func() time.Time) *TokenBucket {
	if capacity < 1 || rate <= 0 {
		panic("invalid bucket")
	}
	return &TokenBucket{}
}
func (*TokenBucket) TryAcquire(permits int) bool { return false }
