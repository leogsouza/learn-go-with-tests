package concurrency

// WebsiteChecker takes a single URL and returns a boolean
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites takes a slice of urls and returns a map of these urls
// checked to a boolean value
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)

	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
