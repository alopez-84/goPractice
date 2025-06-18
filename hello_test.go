package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Albert")
	want := "Hello Albert"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
