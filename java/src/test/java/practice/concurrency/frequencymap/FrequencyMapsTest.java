package practice.concurrency.frequencymap;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.concurrent.*;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;

class FrequencyMapsTest {
    @Test @Timeout(3) void lockedMapCountsUnderContention() throws Exception { assertCounts(new FrequencyMaps.LockedFrequencyMap()); }
    @Test @Timeout(3) void concurrentMapCountsUnderContention() throws Exception { assertCounts(new FrequencyMaps.ConcurrentFrequencyMap()); }
    private static void assertCounts(FrequencyMaps.FrequencyMap map) throws Exception {
        try (var workers = Executors.newFixedThreadPool(8)) {
            for (int i = 0; i < 8; i++) workers.submit(() -> { for (int j = 0; j < 1_000; j++) map.increment("java"); });
            workers.shutdown(); workers.awaitTermination(2, TimeUnit.SECONDS);
        }
        assertEquals(8_000, map.count("java"));
    }
}
