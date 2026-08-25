package practice.concurrency.boundedblockingqueue;

import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;

/** A fixed-capacity, FIFO queue whose operations block when progress is impossible. */
public final class BoundedBlockingQueue<T> {
    private final BlockingQueue<T> queue;

    public BoundedBlockingQueue(int capacity) {
        if (capacity <= 0) {
            throw new IllegalArgumentException("capacity must be positive");
        }
        queue = new ArrayBlockingQueue<>(capacity, true);
    }

    public void put(T value) throws InterruptedException {
        queue.put(value);
    }

    public T take() throws InterruptedException {
        return queue.take();
    }
}
