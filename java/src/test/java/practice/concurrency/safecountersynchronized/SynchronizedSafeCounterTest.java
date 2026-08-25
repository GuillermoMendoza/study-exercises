package practice.concurrency.safecountersynchronized;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;

class SynchronizedSafeCounterTest {
    @Test
    @Timeout(3)
    void preservesEveryIncrementUsingAJavaMonitor() throws InterruptedException {
        var counter = new SynchronizedSafeCounter();
        int workers = 12;
        int incrementsPerWorker = 1_500;
        var ready = new CountDownLatch(workers);
        var start = new CountDownLatch(1);

        try (ExecutorService executor = Executors.newFixedThreadPool(workers)) {
            for (int worker = 0; worker < workers; worker++) {
                executor.submit(() -> {
                    ready.countDown();
                    start.await();
                    for (int increment = 0; increment < incrementsPerWorker; increment++) {
                        counter.increment();
                    }
                    return null;
                });
            }
            ready.await();
            start.countDown();
            executor.shutdown();
            executor.awaitTermination(2, TimeUnit.SECONDS);
        }

        assertEquals(workers * incrementsPerWorker, counter.value());
    }
}
