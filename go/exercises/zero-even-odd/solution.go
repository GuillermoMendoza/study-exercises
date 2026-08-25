package zeroevenodd

// ZeroEvenOdd is LeetCode 1116. Zero, Odd, and Even are called concurrently.
type ZeroEvenOdd struct {
	n        int
	zeroChan chan struct{}
	oddChan  chan struct{}
	evenChan chan struct{}
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	if n < 1 {
		panic("n must be positive")
	}
	z := &ZeroEvenOdd{
		n:        n,
		zeroChan: make(chan struct{}),
		oddChan:  make(chan struct{}),
		evenChan: make(chan struct{}),
	}
	go func() {
		z.zeroChan <- struct{}{}
	}()
	return z
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		<-z.zeroChan
		printNumber(0)
		if i%2 == 0 {
			z.evenChan <- struct{}{}
		} else {
			z.oddChan <- struct{}{}
		}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for i := 2; i <= z.n; i += 2 {
		<-z.evenChan
		printNumber(i)
		if i < z.n {
			z.zeroChan <- struct{}{}
		}
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for i := 1; i <= z.n; i += 2 {
		<-z.oddChan
		printNumber(i)
		if i < z.n {
			z.zeroChan <- struct{}{}
		}
	}
}
