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
		got := Hello("Kamal!", "")
		want := "Hello Kamal!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Kamal!", "Spanish")
		want := "Hola Kamal!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Frenc", func(t *testing.T) {
		got := Hello("Kamal!", "French")
		want := "Bonjour Kamal!"
		assertCorrectMessage(t, got, want)
	})

}
