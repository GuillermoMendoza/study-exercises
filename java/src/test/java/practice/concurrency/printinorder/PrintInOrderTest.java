package practice.concurrency.printinorder;

import static org.junit.jupiter.api.Assertions.assertEquals;
import java.util.*;
import java.util.concurrent.*;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;
class PrintInOrderTest {
 @Test @Timeout(2) void printsFirstSecondThird() throws Exception {
  var printer = new PrintInOrder(); var out = Collections.synchronizedList(new ArrayList<String>());
  try (var e = Executors.newFixedThreadPool(3)) { e.submit(() -> printer.third(() -> out.add("third"))); e.submit(() -> printer.second(() -> out.add("second"))); e.submit(() -> printer.first(() -> out.add("first"))); e.shutdown(); e.awaitTermination(1, TimeUnit.SECONDS); }
  assertEquals(List.of("first", "second", "third"), out);
 }
}
