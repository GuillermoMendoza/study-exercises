# Running average with adders

Use `DoubleAdder` and `LongAdder` to make additions scalable. This exercise guarantees the exact average **after** all writers finish. Unlike the synchronized variant, an average read while writers run does not need to be a single coherent `(sum, count)` snapshot.

Run: `./practice test running-average-atomic`
