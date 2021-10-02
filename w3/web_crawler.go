package main

import (
	"fmt"
	"sync"
)

type SafeCache struct {
	mu sync.Mutex
	v map[string]bool
}

func (c *SafeCache) Store(key string) {
	c.mu.Lock()
	c.v[key] = true
	c.mu.Unlock()
}

func (c *SafeCache) Value(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, 
	safeCache SafeCache, state chan string) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	defer close(state)

	if depth <= 0 || safeCache.Value(url) {
		return
	}

	safeCache.Store(url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	states := make([]chan string, len(urls))
	for i, u := range urls {
		states[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, safeCache, states[i])
	}

	for i := range states {
        for s := range states[i] {
            state <- s
        }
    }
	return
}

func main() {
	safeCache := SafeCache{v: make(map[string]bool)}
	state := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, safeCache, state)

	for s := range state {
        fmt.Println(s)
    }
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
