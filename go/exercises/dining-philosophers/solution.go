package diningphilosophers

// LeetCode 1226. Avoid deadlock while preserving callback order.
type DiningPhilosophers struct{}

func NewDiningPhilosophers() *DiningPhilosophers { return &DiningPhilosophers{} }
func (*DiningPhilosophers) WantsToEat(philosopher int, pickLeft, pickRight, eat, putLeft, putRight func()) {
}
