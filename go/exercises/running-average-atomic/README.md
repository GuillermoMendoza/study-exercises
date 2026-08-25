# Running average with atomics

Implement a CAS loop for the `float64` sum (`math.Float64bits` / `math.Float64frombits`) and an atomic count. After all writers finish, `GetAverage` must be exact. While writers are active, a read may combine a sum and count taken at different instants.

Run: `./practice test running-average-atomic`
