package main

import (
	"fmt"
	"sync"
	"time"
)

// responses := make(map[string]string)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan string, responses *map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, url := range urls {
		// Crawl(url, depth-1, fetcher, ch, responses, wg)
        wg.Add(1)
		go Crawl(url, depth-1, fetcher, ch, responses, wg)

		// wg.Add(1)
		// go func() {
		//     wg.Done()
		//     Crawl(url, depth-1, fetcher, ch, responses, wg)
		// }()
	}
	return
}

func main() {
	start := time.Now()
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	responses := make(map[string]string)
	Crawl("https://golang.org/", 4, fetcher, ch, &responses, &wg)
	wg.Wait()
    // time.Sleep(2 * time.Second)
    // muSToSec := func(s time.)
    delta := time.Now().UnixMicro()-start.UnixMicro()
    fmt.Printf("Program ran: %v ms\n", delta/1000)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    fmt.Println("fetching URL")
    time.Sleep(100 * time.Millisecond)
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
