package practice.concurrency.safecounter;

import java.util.concurrent.atomic.AtomicInteger;

/** A counter that must be safe to share between multiple threads. */
public final class SafeCounter {
    private final AtomicInteger counter = new AtomicInteger(0);

    public void increment() {
        counter.incrementAndGet();
    }

    public int value() {
        return counter.get();
    }
}
