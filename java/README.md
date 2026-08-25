# Java concurrency exercises

This project targets **Java 25**. `./practice` finds the installed JDK 25 on macOS and runs Maven with it, avoiding any other JDK that your global Maven configuration may select.

## Commands

```sh
./practice list
./practice test safe-counter
./practice all
./practice new readers-writers
```

`test` accepts a lowercase kebab-case exercise ID. New exercises are placed in their own package below `src/main/java/practice/concurrency/`, with matching JUnit tests.

The included stubs intentionally fail their acceptance tests. Implement only the production code for an exercise; its test describes the required behavior.

## VS Code

Install the recommended Java extension pack, then open this folder or the repository root. Maven and JUnit tests are discovered automatically. Use the root tasks for a selected exercise or the Testing panel for individual test methods.
