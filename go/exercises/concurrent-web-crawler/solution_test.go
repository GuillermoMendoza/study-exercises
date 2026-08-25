package concurrentwebcrawler

import (
	"context"
	"testing"
	"time"
)

func TestCrawlerDeduplicates(t *testing.T) {
	graph := map[string][]string{"a": {"b", "c", "b"}, "b": {"c"}, "c": {}}
	c := NewConcurrentWebCrawler(2, func(_ context.Context, u string) ([]string, error) { return graph[u], nil })
	r := c.Crawl(context.Background(), "a", time.Second)
	if len(r.Visited) != 3 {
		t.Fatalf("visited=%v, want three URLs", r.Visited)
	}
	if r.PeakConcurrency > 2 {
		t.Fatal("concurrency limit exceeded")
	}
}
