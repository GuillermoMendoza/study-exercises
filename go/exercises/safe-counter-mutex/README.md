# Safe counter with `sync.Mutex`

Implement `MutexSafeCounter` using a plain `int` guarded by `sync.Mutex`. `Increment()` and `Value()` must both lock the same mutex. Do not use `sync/atomic` in this exercise.

Run it with:

```sh
./practice test safe-counter-mutex
```
