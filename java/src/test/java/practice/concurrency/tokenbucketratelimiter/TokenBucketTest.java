package practice.concurrency.tokenbucketratelimiter;
import static org.junit.jupiter.api.Assertions.*; import java.util.concurrent.atomic.AtomicLong; import org.junit.jupiter.api.*;
class TokenBucketTest { @Test void refillsUsingTheInjectedClock(){ var now=new AtomicLong(); var bucket=new TokenBucket(2,2,now::get); assertTrue(bucket.tryAcquire(2)); assertFalse(bucket.tryAcquire(1)); now.addAndGet(500_000_000L); assertTrue(bucket.tryAcquire(1)); } }
