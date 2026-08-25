# Concurrency Practice

Two small, independent workspaces for solving concurrency and LeetCode-style exercises locally:

- [Java](java/README.md) — Java 25, Maven, and JUnit 5.
- [Go](go/README.md) — Go 1.24, standard library, and the race detector.

Open this folder in VS Code. The included tasks can run every exercise or prompt for a single exercise ID. You can also open `java/` or `go/` directly when focusing on one language.

The starter exercises in each language are deliberately incomplete. Their tests compile and run, but fail until you implement the solution. See the full [Java catalogue](java/EXERCISES.md) and [Go catalogue](go/EXERCISES.md).

## Starter exercises

| ID | Focus |
| --- | --- |
| `safe-counter` | Protect a shared counter from lost updates. |
| `safe-counter-synchronized` / `safe-counter-mutex` | Solve the counter again using a Java monitor or Go mutex. |
| `bounded-blocking-queue` | Coordinate producers and consumers with FIFO, capacity, and blocking semantics. |
| `zero-even-odd` | Order three concurrent workers, following LeetCode 1116. |

The catalogue also includes worker pools, frequency maps, crawler simulation, Dining Philosophers, multithreaded FizzBuzz, rate limiting, readers-writer caches, result-stream merging, first-success racing, and ordered parallel maps.

## Fast start

```sh
./java/practice list
./java/practice test safe-counter

./go/practice list
./go/practice test safe-counter
```

Use `./java/practice new my-exercise` or `./go/practice new my-exercise` to create an isolated stub and test from the matching template.
