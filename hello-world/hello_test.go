package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Noah"), "Hello, Noah")
	})
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		assertCorrectMessage(t, Hello(""), "Hello, world")
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
