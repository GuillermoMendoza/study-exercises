package multithreadedfizzbuzz

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestFizzBuzzOrder(t *testing.T) {
	f := NewFizzBuzz(5)
	var out []string
	var mu sync.Mutex
	add := func(s string) { mu.Lock(); defer mu.Unlock(); out = append(out, s) }
	var w sync.WaitGroup
	w.Add(4)
	go func() { defer w.Done(); f.Fizz(func() { add("fizz") }) }()
	go func() { defer w.Done(); f.Buzz(func() { add("buzz") }) }()
	go func() { defer w.Done(); f.FizzBuzz(func() { add("fizzbuzz") }) }()
	go func() { defer w.Done(); f.Number(func(n int) { add(string(rune('0' + n))) }) }()
	done := make(chan struct{})
	go func() { w.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("deadlock")
	}
	if !reflect.DeepEqual(out, []string{"1", "2", "fizz", "4", "buzz"}) {
		t.Fatalf("out=%v", out)
	}
}
