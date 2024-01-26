package main

import (
	"fmt"
	"sync"
)

var fetched = struct {
	m map[string]bool
	sync.Mutex
}{m: make(map[string]bool)}

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	fetched.Lock()
	if fetched.m[url] { // 如果已经抓取过了，就不再抓取了
		fetched.Unlock()
		return
	}
	fetched.m[url] = true // 标记为已抓取
	fetched.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// 可以使用 sync.WaitGroup 来等待一组 goroutine 结束。
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1) // 计数器加 1
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			defer wg.Done() // 计数器减 1
		}(u)
	}
	wg.Wait() // 等待所有 goroutine 结束 （即计数器归 0）
	return
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

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

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fetcher 是填充后的 fakeFetcher。
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
