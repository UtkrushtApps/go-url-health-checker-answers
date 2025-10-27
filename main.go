package main

import (
	"fmt"
	"time"
	"go-url-health-checker/service"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://reddit.com",
		"https://golang.org",
		"https://example.com",
		"https://amazon.com",
		"https://microsoft.com",
		"https://apple.com",
		"https://twitter.com",
	}

	start := time.Now()
	results := service.CheckURLs(urls)

	for _, result := range results {
		if result.Err != nil {
			fmt.Printf("[%-25s] ERROR: %v\n", result.URL, result.Err)
		} else {
			fmt.Printf("[%-25s] STATUS: %s\n", result.URL, result.Status)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("\nTotal Time Taken: %v\n", elapsed)
}
