package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type FetchState struct {
	mu sync.Mutex
	active sync.WaitGroup
	urls map[string]string
}

func (url_map *FetchState) contains(url string) bool{
	url_map.mu.Lock()
	defer url_map.mu.Unlock()

	_, ok := url_map.urls[url]
	return ok
}

func (url_map *FetchState) add(url string, body string) {
	url_map.mu.Lock()
	defer url_map.mu.Unlock()

	url_map.urls[url] = body
}

func (state *FetchState) start_fetch() {
	state.active.Add(1)
}

func (state *FetchState) finish_fetch() {
	state.active.Done()
}

func (state *FetchState) wait() {
	state.active.Wait()
}



// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, dups *FetchState) {
	defer dups.finish_fetch()

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	dups.add(url, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		if !dups.contains(u) {
			dups.start_fetch()
			go Crawl(u, depth-1, fetcher, dups)
		}
	}
}

func main() {
	dups_map := FetchState{urls: make(map[string]string) }

	dups_map.start_fetch()
	Crawl("https://golang.org/", 4, fetcher, &dups_map)
	dups_map.wait()
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
	"https://golang.org/cmd/": &fakeResult{
		"Commands",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/nutties",
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
