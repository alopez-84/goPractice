package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Albert")
		want := "Hello, Albert"

		assertCorrectMessage(t, got, want)
	})

	t.Run("testing with an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Want %q, got %q", got, want)
	}
}
