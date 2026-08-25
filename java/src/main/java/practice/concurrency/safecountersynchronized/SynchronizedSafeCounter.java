package practice.concurrency.safecountersynchronized;

/**
 * A counter that must use Java's synchronized monitor mechanism for all access
 * to its mutable state. Do not use an atomic variable for this exercise.
 */
public final class SynchronizedSafeCounter {
    private int value = 0;

    public synchronized void increment() {
        value++;
    }

    public synchronized int value() {
        return value;
    }
}
