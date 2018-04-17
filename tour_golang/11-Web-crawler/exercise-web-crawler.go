package main

import (
	"fmt"
	"sync"
)

// Fetcher - Fetch returns the body of URL and a slice of URLs found on that page
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// URLCache -  Cache struct for urls
type URLCache struct {
	urls map[string]bool
	mux  sync.Mutex
}

// Set - add URL to cache
func (cache *URLCache) Set(url string) {
	cache.mux.Lock()
	cache.urls[url] = true
	cache.mux.Unlock()
}

// Contains - check if cache contains URL
func (cache *URLCache) Contains(url string) bool {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	return cache.urls[url]
}

// Crawler - cache, fetcher, waitGroup
type Crawler struct {
	urlCache  URLCache
	fetcher   Fetcher
	waitGroup sync.WaitGroup
}

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth
func (crawler *Crawler) Crawl(url string, depth int) {
	defer crawler.waitGroup.Done()

	if depth <= 0 {
		return
	}

	if crawler.urlCache.Contains(url) {
		return
	}

	_, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Printf("Crawl: error when fetching URL %s %v\n", url, err)

		return
	}

	fmt.Printf("Visited %s - %d URLs found\n", url, len(urls))

	crawler.urlCache.Set(url)

	for _, u := range urls {
		crawler.waitGroup.Add(1)
		go crawler.Crawl(u, depth-1)
	}
}

func main() {
	crawler := Crawler{
		urlCache: URLCache{urls: make(map[string]bool)},
		fetcher:  fetcher,
	}

	crawler.waitGroup = sync.WaitGroup{}
	crawler.waitGroup.Add(1)

	go crawler.Crawl("http://golang.org/", 4)
	crawler.waitGroup.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
