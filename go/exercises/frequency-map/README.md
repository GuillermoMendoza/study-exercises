# Frequency map race progression

First implement `UnsafeFrequencyMap` with an ordinary map and run a concurrent workload under `go test -race` to observe the race. Then implement `LockedFrequencyMap` with `sync.Mutex` and `ConcurrentFrequencyMap` with `sync.Map` or a concurrent map plus atomic counters. The normal tests validate only the two safe variants.
