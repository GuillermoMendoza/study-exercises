package practice.concurrency.slidingwindowcounter;

import java.time.Instant;
import java.util.ArrayDeque;
import java.util.Deque;

public final class SlidingWindowCounterComplex {
    private static final long WINDOW_SECONDS = 5* 60;
    private final Deque<Bucket> buckets = new ArrayDeque<>();
    private int totalCalls = 0;

    public synchronized void increment() {
        long now = Instant.now().getEpochSecond();
        removeExpired(Instant.now().getEpochSecond());
        
        Bucket latest = buckets.peekLast();

        if (latest != null && latest.timestamp == now) {
            latest.count++;
        } else {
            buckets.addLast(new Bucket(now, 1));
        }

        totalCalls++;
    }

    public synchronized int getValue() {
        removeExpired(Instant.now().getEpochSecond());
        return totalCalls;
    } 

    private void removeExpired(long now) {
        long cutoff = now - WINDOW_SECONDS;

        while (!buckets.isEmpty() && buckets.peekFirst().timestamp <= cutoff) {
            totalCalls -= buckets.removeFirst().count;
        }
    }

    private static final class Bucket {
        final long timestamp;
        int count;

        Bucket(long timestamp, int count) {
            this.timestamp = timestamp;
            this.count = count;
        }
    }
}
