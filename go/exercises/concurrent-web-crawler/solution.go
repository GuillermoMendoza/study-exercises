package concurrentwebcrawler

import (
	"context"
	"time"
)

type Fetcher func(context.Context, string) ([]string, error)
type CrawlResult struct {
	Visited         map[string]struct{}
	Failed          map[string]error
	TimedOut        bool
	PeakConcurrency int
}
type ConcurrentWebCrawler struct{}

func NewConcurrentWebCrawler(max int, fetch Fetcher) *ConcurrentWebCrawler {
	if max < 1 {
		panic("max must be positive")
	}
	return &ConcurrentWebCrawler{}
}
func (*ConcurrentWebCrawler) Crawl(ctx context.Context, seed string, timeout time.Duration) CrawlResult {
	return CrawlResult{Visited: map[string]struct{}{}, Failed: map[string]error{}}
}
