package practice.concurrency.fixedworkerpool;

import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.Callable;
import java.util.concurrent.Future;
import java.util.concurrent.FutureTask;
import java.util.concurrent.LinkedBlockingDeque;
import java.util.concurrent.RejectedExecutionException;

/** Implement a fixed-size pool without delegating scheduling to ExecutorService. */
public final class FixedWorkerPool implements AutoCloseable {
    private final BlockingQueue<Runnable> taskQueue = new LinkedBlockingDeque<>();
    private final List<Thread> workers;
    private volatile boolean triggerShutdown = false;

    public FixedWorkerPool(int workerCount) {
        if (workerCount < 1) {
            throw new IllegalArgumentException("workerCount must be positive");
        }
        workers = new ArrayList<>(workerCount);
        for (int i = 0; i < workerCount; i++) {
            Thread worker = new Thread(new WorkerLoop());
            workers.add(worker);
            worker.start();
        }
    }

    public <T> Future<TaskResult<T>> submit(Callable<T> task) {
        if (task == null) {
            throw new NullPointerException("Task cannot be null");
        }

        Callable<TaskResult<T>> wrappedTask = () -> {
            try {
                T value = task.call();
                return new TaskResult<>(value, null);
            } catch (Throwable ex) {
                return new TaskResult<>(null, ex);
            }
        };

        FutureTask<TaskResult<T>> futureTask = new FutureTask<>(wrappedTask);

        synchronized (this) {
            if (triggerShutdown) {
                throw new RejectedExecutionException("Pool was shutdown");
            }
            taskQueue.add(futureTask);
        }

        return futureTask;
    }

    public void shutdown() { 
        synchronized (this) {
            if (triggerShutdown) {
                return;
            }
            triggerShutdown = true;

            for (Thread worker : workers) {
                worker.interrupt();
            }
        }
    }

    public boolean awaitTermination(Duration timeout) throws InterruptedException { 
        long deadline = System.nanoTime() + timeout.toNanos();

        for (Thread worker : workers) {
            long remainingNanos = deadline - System.nanoTime();
            if (remainingNanos <= 0) {
                return false;
            }
            worker.join(remainingNanos / 1_000_000, (int) (remainingNanos % 1_000_000));
            if (worker.isAlive()) {
                return false;
            }
        }
        return true;
    }

    @Override
    public void close() { 
        shutdown(); 
    }

    private final class WorkerLoop implements Runnable {
        @Override
        public void run() {
            try {
                while (!triggerShutdown || !taskQueue.isEmpty()) {
                    try {
                        Runnable task = taskQueue.take();
                        task.run();
                    } catch (InterruptedException ex) {
                        if (triggerShutdown && taskQueue.isEmpty()) {
                            break;
                        }
                    }
                }
            } catch (Throwable ex) {
                ex.printStackTrace();
            }
        }
    }
}
