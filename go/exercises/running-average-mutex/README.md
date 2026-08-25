# Running average with `sync.Mutex`

The original two-field implementation has races in both `AddNumber` and `GetAverage`. Add one `sync.Mutex` that protects the sum/count pair in both methods. A concurrent reader must see a coherent average.

Run: `./practice test running-average-mutex`
