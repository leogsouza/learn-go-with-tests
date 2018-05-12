package concurrency

import "time"

// WebsiteChecker takes a single URL and returns a boolean
type WebsiteChecker func(string) bool

// CheckWebsites takes a slice of urls and returns a map of these urls
// checked to a boolean value
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)

	}

	time.Sleep(2 * time.Second)

	return results
}
