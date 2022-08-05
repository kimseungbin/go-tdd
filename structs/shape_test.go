package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 2.0)
	want := 24.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
