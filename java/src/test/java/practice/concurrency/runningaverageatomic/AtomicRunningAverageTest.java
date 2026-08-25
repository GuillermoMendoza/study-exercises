package practice.concurrency.runningaverageatomic;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;

class AtomicRunningAverageTest {
    private static final double DELTA = 1.0e-9;

    @Test
    void returnsZeroBeforeAnyValuesAreAdded() {
        assertEquals(0.0, new AtomicRunningAverage().getAverage(), DELTA);
    }

    @Test
    @Timeout(3)
    void returnsTheExactAverageAfterConcurrentWritersFinish() throws InterruptedException {
        var average = new AtomicRunningAverage();
        int writers = 12;
        int additionsPerWriter = 2_000;
        var ready = new CountDownLatch(writers);
        var start = new CountDownLatch(1);

        try (ExecutorService executor = Executors.newFixedThreadPool(writers)) {
            for (int writer = 0; writer < writers; writer++) {
                double value = writer;
                executor.submit(() -> {
                    ready.countDown();
                    start.await();
                    for (int addition = 0; addition < additionsPerWriter; addition++) {
                        average.addNumber(value);
                    }
                    return null;
                });
            }
            ready.await();
            start.countDown();
            executor.shutdown();
            executor.awaitTermination(2, TimeUnit.SECONDS);
        }

        assertEquals((writers - 1) / 2.0, average.getAverage(), DELTA);
    }
}
