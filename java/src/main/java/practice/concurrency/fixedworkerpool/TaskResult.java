package practice.concurrency.fixedworkerpool;

/** Outcome of one accepted task. Exactly one of value and failure is meaningful. */
public record TaskResult<T>(T value, Throwable failure) {
    public boolean succeeded() { return failure == null; }
}
