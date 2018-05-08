package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Cris", "")
		want := "Hello, Cris"

		assertCorrectMessage(t, got, want)

	})

	t.Run("empty string defaults to 'Hello, World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	
	t.Run("in Spanish", func(t *testing.T) {
	    got := Hello("Eloide", "Spanish")
	    want := "Hola, Eloide"
	    
	    assertCorrectMessage(t, got, want)
	})
	
	t.Run("in French", func(t *testing.T) {
	    got := Hello("Jean", "French")
	    want := "Bonjour, Jean"
	    
	    assertCorrectMessage(t, got, want)
	})

}
