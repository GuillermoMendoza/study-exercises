package practice.concurrency.fixedworkerpool;

import static org.junit.jupiter.api.Assertions.*;

import java.time.Duration;
import java.util.concurrent.RejectedExecutionException;
import java.util.concurrent.TimeUnit;
import org.junit.jupiter.api.Test;

class FixedWorkerPoolTest {
    @Test
    void collectsSuccessAndFailureAndTerminates() throws Exception {
        try (var pool = new FixedWorkerPool(2)) {
            var success = pool.submit(() -> 42);
            var failure = pool.submit(() -> { throw new IllegalStateException("boom"); });
            assertEquals(42, success.get(1, TimeUnit.SECONDS).value());
            assertTrue(failure.get(1, TimeUnit.SECONDS).failure() instanceof IllegalStateException);
            pool.shutdown();
            assertTrue(pool.awaitTermination(Duration.ofSeconds(1)));
            assertThrows(RejectedExecutionException.class, () -> pool.submit(() -> 1));
        }
    }
}
