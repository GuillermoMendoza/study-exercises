package runningaverageatomic

import (
	"math"
	"sync/atomic"
)

// AtomicRunningAverage should use a CAS loop over sumBits plus an atomic count.
// Its exact-average guarantee applies after all writers have finished.
type AtomicRunningAverage struct {
	sumBits atomic.Uint64
	count   atomic.Uint64
}

func NewAtomicRunningAverage() *AtomicRunningAverage {
	return &AtomicRunningAverage{}
}

func (r *AtomicRunningAverage) AddNumber(number float64) {
	for {
		currentBits := r.sumBits.Load()
		currentSum := math.Float64frombits(currentBits) + number
		newBits := math.Float64bits(currentSum)
		if r.sumBits.CompareAndSwap(currentBits, newBits) {
			break
		}
	}

	r.count.Add(1)
}

func (r *AtomicRunningAverage) GetAverage() float64 {
	totalElements := r.count.Load()
	if totalElements == 0 {
		return 0.0
	}

	totalSum := math.Float64frombits(r.sumBits.Load())
	return totalSum / float64(totalElements)
}
