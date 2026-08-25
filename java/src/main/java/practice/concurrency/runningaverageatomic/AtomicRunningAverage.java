package practice.concurrency.runningaverageatomic;

import java.util.concurrent.atomic.DoubleAdder;
import java.util.concurrent.atomic.LongAdder;

/**
 * Implement an eventually consistent running average. After all writers have
 * completed, getAverage must be exact; a concurrent read may combine adder
 * snapshots from different instants.
 */
public final class AtomicRunningAverage {
    private final DoubleAdder sum = new DoubleAdder();
    private final LongAdder count = new LongAdder();

    public void addNumber(double number) {
        sum.add(number);
        count.increment();
    }

    public double getAverage() {
        long totalElements = count.sum();
        if(totalElements == 0) {
            return 0.0;
        }

        return sum.sum()/totalElements;
    }
}
