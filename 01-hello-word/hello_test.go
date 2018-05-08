package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Cris")
	want := "Hello, Cris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
