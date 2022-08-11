package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(tb testing.TB, got, want string) {
	tb.Helper()

	if got != want {
		tb.Errorf("got %q want %q", got, want)
	}
}
