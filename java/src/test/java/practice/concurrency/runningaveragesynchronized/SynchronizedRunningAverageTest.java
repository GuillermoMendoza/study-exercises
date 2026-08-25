package practice.concurrency.runningaveragesynchronized;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.util.concurrent.ConcurrentLinkedQueue;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;

class SynchronizedRunningAverageTest {
    private static final double DELTA = 1.0e-9;

    @Test
    void returnsZeroBeforeAnyValuesAreAdded() {
        assertEquals(0.0, new SynchronizedRunningAverage().getAverage(), DELTA);
    }

    @Test
    @Timeout(4)
    void keepsAConsistentAverageWhileWorkersAddValues() throws InterruptedException {
        var average = new SynchronizedRunningAverage();
        int writers = 12;
        int additionsPerWriter = 2_000;
        var ready = new CountDownLatch(writers + 1);
        var start = new CountDownLatch(1);
        var writersDone = new CountDownLatch(writers);
        var writing = new AtomicBoolean(true);
        var observed = new ConcurrentLinkedQueue<Double>();

        try (ExecutorService executor = Executors.newFixedThreadPool(writers + 1)) {
            executor.submit(() -> {
                ready.countDown();
                start.await();
                while (writing.get()) {
                    observed.add(average.getAverage());
                }
                return null;
            });
            for (int writer = 0; writer < writers; writer++) {
                double value = writer;
                executor.submit(() -> {
                    ready.countDown();
                    start.await();
                    try {
                        for (int addition = 0; addition < additionsPerWriter; addition++) {
                            average.addNumber(value);
                        }
                    } finally {
                        writersDone.countDown();
                    }
                    return null;
                });
            }
            ready.await();
            start.countDown();
            writersDone.await();
            writing.set(false);
            executor.shutdown();
            executor.awaitTermination(3, TimeUnit.SECONDS);
        } finally {
            writing.set(false);
        }

        for (double snapshot : observed) {
            assertTrue(Double.isFinite(snapshot));
            assertTrue(snapshot >= 0.0 && snapshot <= writers - 1.0);
        }
        assertEquals((writers - 1) / 2.0, average.getAverage(), DELTA);
    }
}
