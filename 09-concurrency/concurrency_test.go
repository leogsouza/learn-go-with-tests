package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "https://leogsouza.info" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://github.com",
		"https://twitter.com",
		"https://leogsouza.info",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	expectedResults := map[string]bool{
		"https://github.com":     true,
		"https://twitter.com":    true,
		"https://leogsouza.info": false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}
