# Additional concurrency exercises

All exercises are intentionally unsolved. Their tests are the contract.

- `fixed-worker-pool`: implement a bounded worker pool; record success/failure per accepted task and shut down cleanly.
- `frequency-map`: first observe the unsafe map race, then implement monitor-locked and concurrent-map variants.
- `print-in-order`: LeetCode 1114; coordinate `first`, `second`, and `third`.
- `concurrent-web-crawler`: crawl an injected in-memory graph with bounded fetching, deduplication, cancellation, and a deadline.
- `dining-philosophers`: LeetCode 1226; avoid deadlock and adjacent philosophers eating together.
- `multithreaded-fizzbuzz`: LeetCode 1195 callback ordering.
- `token-bucket-rate-limiter`: a thread-safe bucket using the injected monotonic clock.
- `readers-writer-cache`: use `ReentrantReadWriteLock` semantics and load a missing key only once.
- `merge-result-streams`: merge completed queue-backed streams while retaining each source's ordering.
- `first-successful-result`: return the first successful task, ignore failures, and cancel the rest.
- `parallel-map-ordered`: map with bounded concurrency, retain input order, and fail fast on mapper error.
- `running-average-synchronized`: protect a `double` sum and count as one coherent snapshot with `synchronized`.
- `running-average-atomic`: use adders for a correct final average after writers finish; intermediate reads need not be a coherent pair.
