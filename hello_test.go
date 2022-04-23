package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(tb testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %Q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
