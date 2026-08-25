package practice.concurrency.concurrentwebcrawler;

import java.time.Duration;
import java.util.*;

/** Crawl only through the supplied fetcher; never perform real network I/O. */
public final class ConcurrentWebCrawler {
    @FunctionalInterface public interface Fetcher { List<String> fetch(String url) throws Exception; }
    public record CrawlResult(Set<String> visited, Set<String> failed, boolean timedOut, int peakConcurrency) { }
    public ConcurrentWebCrawler(int maxConcurrency, Fetcher fetcher) { if (maxConcurrency < 1) throw new IllegalArgumentException(); }
    public CrawlResult crawl(String seed, Duration timeout) { return new CrawlResult(Set.of(), Set.of(), false, 0); }
}
