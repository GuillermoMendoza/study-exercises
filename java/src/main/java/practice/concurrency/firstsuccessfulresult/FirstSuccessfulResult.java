package practice.concurrency.firstsuccessfulresult;

import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;
import java.util.concurrent.Callable;
import java.util.concurrent.CancellationException;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorCompletionService;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;

/** Runs attempts concurrently and returns the first successful result. */
public final class FirstSuccessfulResult {
    private FirstSuccessfulResult() { }

    public static <T> T get(List<? extends Callable<T>> attempts, Duration timeout) throws Exception {
        Objects.requireNonNull(attempts, "attempts");
        Objects.requireNonNull(timeout, "timeout");
        
        if (attempts.isEmpty()) {
            throw new IllegalArgumentException("No attempts provided");
        }
        
        if (timeout.isNegative()) {
            throw new IllegalArgumentException("timeout must not be negative");
        }

        ExecutorService executor = Executors.newVirtualThreadPerTaskExecutor();
        var completions = new ExecutorCompletionService<T>(executor);
        var futures = new ArrayList<Future<T>>(attempts.size());
        long deadline = System.nanoTime() + timeout.toNanos();
        Throwable firstFailure = null;

        try {
            for (Callable<T> attempt : attempts) {
                futures.add(completions.submit(Objects.requireNonNull(attempt, "attempt")));
            }

            for (int completedCount = 0; completedCount < attempts.size(); completedCount++) {
                long remainingNanos = deadline - System.nanoTime();
                if (remainingNanos <= 0) {
                    throw new TimeoutException("No attempt succeeded before the deadline");
                }

                Future<T> completed = completions.poll(remainingNanos, TimeUnit.NANOSECONDS);
                if (completed == null) {
                    throw new TimeoutException("No attempt succeeded before the deadline");
                }

                try {
                    return completed.get();
                } catch (ExecutionException exception) {
                    if (firstFailure == null) {
                        firstFailure = exception.getCause();
                    }
                } catch (CancellationException exception) {
                    if (firstFailure == null) {
                        firstFailure = exception;
                    }
                }
            }

            throw new Exception("All attempts failed", firstFailure);
        } finally {
            for (Future<T> future : futures) {
                future.cancel(true);
            }
            executor.shutdownNow();
        }
    }
}
