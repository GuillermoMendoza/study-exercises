package practice.concurrency.concurrentwebcrawler;
import static org.junit.jupiter.api.Assertions.*;
import java.time.Duration; import java.util.*; import org.junit.jupiter.api.Test;
class ConcurrentWebCrawlerTest {
 @Test void deduplicatesAndHonorsTheLimit() {
  var graph = Map.of("a", List.of("b", "c", "b"), "b", List.of("c"), "c", List.<String>of());
  var crawler = new ConcurrentWebCrawler(2, url -> graph.getOrDefault(url, List.of()));
  var result = crawler.crawl("a", Duration.ofSeconds(1));
  assertEquals(Set.of("a", "b", "c"), result.visited()); assertTrue(result.peakConcurrency() <= 2);
 }
}
