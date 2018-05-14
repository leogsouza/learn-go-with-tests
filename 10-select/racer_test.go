package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://www.coursera.org"
	fastURL := "https://www.github.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
