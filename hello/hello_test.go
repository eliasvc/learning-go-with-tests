package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := HelloWorld("Elias", "")
		want := "Hello, Elias"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string results in 'World'", func(t *testing.T) {
		got := HelloWorld("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := HelloWorld("Elias", "Spanish")
		want := "Hola, Elias"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := HelloWorld("Elias", "French")
		want := "Bonjour, Elias"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Italian", func(t *testing.T) {
		got := HelloWorld("Elias", "Italian")
		want := "Ciao, Elias"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got string, want string) {
	// t.Helper() lets testing know that this function is a helper so if something fails
	// the error message will point back to the line of code in the test instead
	// of here.
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
