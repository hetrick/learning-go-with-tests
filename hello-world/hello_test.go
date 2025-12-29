package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello defaults", func(t *testing.T) {
		assertCorrectMessage(t, Hello("", ""), "Hello, world")
	})
	t.Run("hello in English", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Noah", "English"), "Hello, Noah")
	})
	t.Run("hello in Spanish", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Noah", "Spanish"), "Hola, Noah")
	})
	t.Run("hello in French", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Noah", "French"), "Bonjour, Noah")
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
