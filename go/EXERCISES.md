# Additional concurrency exercises

These match the Java exercise IDs and are intentionally unsolved. Run one with `./practice test <id>`; Go always enables the race detector.

The crawler uses an injected local fetch function, the rate limiter uses an injected clock, and timeout/failure exercises use `context.Context`. The `frequency-map` README/test setup includes a separate race-demo stage so the deliberately unsafe implementation does not poison normal practice runs.

- `running-average-mutex`: protect the sum/count pair with `sync.Mutex` so every read is coherent.
- `running-average-atomic`: use atomic sum/count state for a correct final average after writers finish; intermediate reads may combine snapshots.
