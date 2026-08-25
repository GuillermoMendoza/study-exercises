package practice.concurrency.diningphilosophers;
import static org.junit.jupiter.api.Assertions.*; import java.util.concurrent.*; import java.util.concurrent.atomic.*; import org.junit.jupiter.api.*;
class DiningPhilosophersTest {
 @Test @Timeout(3) void everyoneEventuallyEats() throws Exception { var table = new DiningPhilosophers(); var eaten = new AtomicInteger(); try (var e = Executors.newFixedThreadPool(5)) { for (int p=0;p<5;p++) { int id=p; e.submit(() -> table.wantsToEat(id, ()->{}, ()->{}, eaten::incrementAndGet, ()->{}, ()->{})); } e.shutdown(); assertTrue(e.awaitTermination(2, TimeUnit.SECONDS)); } assertEquals(5, eaten.get()); }
}
