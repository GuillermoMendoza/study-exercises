package practice.concurrency.tokenbucketratelimiter;
import java.util.function.LongSupplier;
/** Use the supplied monotonic nanosecond clock; synchronize refill and acquisition. */
public final class TokenBucket {
 public TokenBucket(int capacity, double tokensPerSecond, LongSupplier nanoTime) { if(capacity<1||tokensPerSecond<=0) throw new IllegalArgumentException(); }
 public boolean tryAcquire(int permits) { return false; }
}
