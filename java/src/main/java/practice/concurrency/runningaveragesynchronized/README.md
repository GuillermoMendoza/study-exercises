# Running average with `synchronized`

The original `sum += number` and `count++` update two shared fields without a single critical section. Implement `addNumber` and `getAverage` with the same Java monitor so readers see a coherent `(sum, count)` snapshot while workers are adding values.

Run: `./practice test running-average-synchronized`
