package practice.concurrency.zeroevenodd;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;

class ZeroEvenOddTest {
    @Test
    void rejectsNonPositiveN() {
        org.junit.jupiter.api.Assertions.assertThrows(IllegalArgumentException.class, () -> new ZeroEvenOdd(0));
    }

    @Test
    @Timeout(3)
    void emitsZeroThenEachNumberInOrder() throws InterruptedException {
        var exercise = new ZeroEvenOdd(5);
        List<Integer> output = java.util.Collections.synchronizedList(new ArrayList<>());

        try (ExecutorService executor = Executors.newFixedThreadPool(3)) {
            executor.submit(() -> runInterruptibly(() -> exercise.zero(output::add)));
            executor.submit(() -> runInterruptibly(() -> exercise.even(output::add)));
            executor.submit(() -> runInterruptibly(() -> exercise.odd(output::add)));
            executor.shutdown();
            org.junit.jupiter.api.Assertions.assertTrue(executor.awaitTermination(2, TimeUnit.SECONDS));
        }

        assertEquals(List.of(0, 1, 0, 2, 0, 3, 0, 4, 0, 5), output);
    }

    private static void runInterruptibly(InterruptibleAction action) {
        try {
            action.run();
        } catch (InterruptedException exception) {
            Thread.currentThread().interrupt();
        }
    }

    @FunctionalInterface
    private interface InterruptibleAction {
        void run() throws InterruptedException;
    }
}
