package practice.concurrency.slidingwindowcounter;

import java.time.Instant;
import java.util.ArrayDeque;
import java.util.Deque;

public final class SlidingWindowCounterEasy {
    private static final long WINDOW_SECONDS = 5 * 60;
    private final Deque<Long> timestamps = new ArrayDeque<>();

    public synchronized void increment() {
        long now = Instant.now().getEpochSecond();
        removeExpired(now);
        timestamps.addLast(now);
    }

    public synchronized int getValue() {
        long now = Instant.now().getEpochSecond();
        removeExpired(now);
        return timestamps.size();
    }

    private void removeExpired(long now) {
        long cutoff = now - WINDOW_SECONDS;

        while (!timestamps.isEmpty() && timestamps.peekFirst() <= cutoff) {
            timestamps.removeFirst();
        }
    }
}
