# Go concurrency exercises

This module targets the installed Go **1.24.4** (the Go 1.24 series). It has no external dependencies. Every test runs with the race detector so unsafely shared state is reported while you practice.

## Commands

```sh
./practice list
./practice test safe-counter
./practice all
./practice new readers-writers
```

Each exercise lives in `exercises/<exercise-id>/`. New exercises contain a compilable `Solution` stub and a deliberately failing test.

The existing exercise APIs use idiomatic Go names, while their contracts match the Java counterparts. Use the Testing panel in VS Code or the root tasks to run them.
