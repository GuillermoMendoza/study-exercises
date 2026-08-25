package multithreadedfizzbuzz

// LeetCode 1195.
type FizzBuzz struct{}

func NewFizzBuzz(n int) *FizzBuzz {
	if n < 1 {
		panic("n must be positive")
	}
	return &FizzBuzz{}
}
func (*FizzBuzz) Fizz(func())      {}
func (*FizzBuzz) Buzz(func())      {}
func (*FizzBuzz) FizzBuzz(func())  {}
func (*FizzBuzz) Number(func(int)) {}
