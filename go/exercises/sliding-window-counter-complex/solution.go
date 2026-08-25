package slidingwindowcountercomplex

import (
	"sync"
	"time"
)

// 5 minutes to seconds
const windowSeconds int64 = 5 * 60

type Bucket struct {
	timestamp int64
	count     int64
}

type SlidingWindowCounter struct {
	mu         sync.Mutex
	buckets    []Bucket
	totalCalls int64
}

func (s *SlidingWindowCounter) Increment() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().Unix()
	s.removeExpired(now)

	newestBucketIndex := len(s.buckets) - 1

	if newestBucketIndex >= 0 && s.buckets[newestBucketIndex].timestamp == now {
		s.buckets[newestBucketIndex].count++
	} else {
		s.buckets = append(s.buckets, Bucket{
			timestamp: now,
			count:     1,
		})
	}

	s.totalCalls++
}

func (s *SlidingWindowCounter) GetValue() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.removeExpired(time.Now().Unix())

	return s.totalCalls
}

func (s *SlidingWindowCounter) removeExpired(now int64) {
	cutoff := now - windowSeconds
	expiredBuckets := 0

	for expiredBuckets < len(s.buckets) && s.buckets[expiredBuckets].timestamp <= cutoff {
		s.totalCalls -= s.buckets[expiredBuckets].count
		expiredBuckets++
	}

	s.buckets = s.buckets[expiredBuckets:]
}
