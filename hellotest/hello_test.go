package main

import "testing"

func TestHello(t *testing.T) {
	correctMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		correctMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		correctMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		correctMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Sophie", "French")
		want := "Bonjour, Sophie"
		correctMessage(t, got, want)
	})
}
