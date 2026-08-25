package printinorder

// LeetCode 1114. Calls arrive concurrently and callbacks must run first, second, third.
type PrintInOrder struct{}

func NewPrintInOrder() *PrintInOrder   { return &PrintInOrder{} }
func (*PrintInOrder) First(fn func())  {}
func (*PrintInOrder) Second(fn func()) {}
func (*PrintInOrder) Third(fn func())  {}
