package main

import (
	"fmt"
	"net/http"
)

type WebsiteChecker func(string) bool

func realWebsiteChecker(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

func main() {
	urls := []string{"https://google.com", "https://github.com", "https://foo.org"}
	results := CheckWebsites(realWebsiteChecker, urls)
	for k, v := range results {
		fmt.Printf("%s: %v\n", k, v)
	}
}
