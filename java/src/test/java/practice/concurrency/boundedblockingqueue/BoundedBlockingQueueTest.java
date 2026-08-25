package practice.concurrency.boundedblockingqueue;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertThrows;

import java.time.Duration;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;

class BoundedBlockingQueueTest {
    @Test
    void rejectsNonPositiveCapacity() {
        assertThrows(IllegalArgumentException.class, () -> new BoundedBlockingQueue<Integer>(0));
    }

    @Test
    @Timeout(3)
    void keepsFifoOrderAndUnblocksAProducerAfterTake() throws Exception {
        var queue = new BoundedBlockingQueue<Integer>(1);
        queue.put(1);

        try (ExecutorService executor = Executors.newSingleThreadExecutor()) {
            Future<?> blockedPut = executor.submit(() -> {
                queue.put(2);
                return null;
            });

            Thread.sleep(Duration.ofMillis(100));
            assertFalse(blockedPut.isDone(), "put must wait while the queue is full");
            assertEquals(1, queue.take());
            blockedPut.get(1, TimeUnit.SECONDS);
            assertEquals(2, queue.take());
        }
    }
}
