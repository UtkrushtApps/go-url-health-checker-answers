package service

import (
	"go-url-health-checker/handler"
	"sync"
	"time"
)

// CheckURLs checks the health of each URL concurrently and streams the results.
func CheckURLs(urls []string) []handler.Result {
	resultsCh := make(chan handler.Result, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			result := handler.CheckURL(u)
			resultsCh <- result
		}(url)
	}

	// Close resultsCh after all goroutines complete
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	results := make([]handler.Result, 0, len(urls))
	for res := range resultsCh {
		results = append(results, res)
	}

	return results
}
