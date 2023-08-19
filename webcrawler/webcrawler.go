package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, mu *sync.Mutex, responses map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
    if _, ok := responses[url]; ok {
        fmt.Printf("Already fetched URL %v, skipping\n", url)
        return
    }
	// TODO: Don't fetch the same URL twice.
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)

    mu.Lock()
    responses[url] = body
    mu.Unlock()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, url := range urls {
        wg.Add(1)
		go Crawl(url, depth-1, fetcher, mu, responses, wg)
	}
	return
}

func main() {
	start := time.Now()

	// ch := make(chan string)
    mu := sync.Mutex{}

	var wg sync.WaitGroup

	wg.Add(1)
	responses := make(map[string]string)
	Crawl("https://golang.org/", 4, fetcher, &mu, responses, &wg)
	wg.Wait()

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
