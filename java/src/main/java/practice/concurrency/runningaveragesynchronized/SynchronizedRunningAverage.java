package practice.concurrency.runningaveragesynchronized;

/**
 * Start from this deliberately unsafe two-field implementation. Make both
 * methods synchronize on the same monitor so a read observes one complete
 * (sum, count) state.
 */
public final class SynchronizedRunningAverage {
    private double sum = 0.0;
    private int count = 0;

    public synchronized void addNumber(double number) {
        sum += number;
        count++;
    }

    public synchronized double getAverage() {
        if (count == 0) {
            return 0.0;
        }
        return sum / count;
    }
}
