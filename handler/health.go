package handler

import (
	"errors"
	"fmt"
	"time"
)

type Result struct {
	URL    string
	Status string
	Err    error
}

// Simulate URL health: hardcoded delays & errors for a few URLs to mimic real-world behavior
func CheckURL(url string) Result {
	// Simulated fake responses for demonstration
	// (delays and timeouts)
	simulatedDelays := map[string]time.Duration{
		"https://google.com":         200 * time.Millisecond,
		"https://github.com":         300 * time.Millisecond,
		"https://stackoverflow.com":  700 * time.Millisecond,
		"https://reddit.com":         1200 * time.Millisecond, // will timeout
		"https://golang.org":         250 * time.Millisecond,
		"https://example.com":        400 * time.Millisecond,
		"https://amazon.com":         1700 * time.Millisecond, // will timeout
		"https://microsoft.com":      800 * time.Millisecond,
		"https://apple.com":          200 * time.Millisecond,
		"https://twitter.com":        2100 * time.Millisecond, // will timeout
	}
	timeout := 1000 * time.Millisecond
	delay, ok := simulatedDelays[url]
	if !ok {
		delay = 300 * time.Millisecond
	}

	// Channel to simulate work
	done := make(chan struct{})
	var status string
	var statusErr error
	go func() {
		time.Sleep(delay) // simulate network delay
		status = "Online"
		statusErr = nil
		done <- struct{}{}
	}()

	select {
	case <-done:
		return Result{URL: url, Status: status, Err: statusErr}
	case <-time.After(timeout):
		return Result{URL: url, Status: "Timeout", Err: errors.New(fmt.Sprintf("timeout after %v", timeout))}
	}
}
