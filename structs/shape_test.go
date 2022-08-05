package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 2.0)
	want := 24.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func ExamplePerimeter() {
	perimeter := Perimeter(5.0, 5.0)
	fmt.Println(perimeter)
	// Output: 20
}

func TestArea(t *testing.T) {
	got := Area(2.0, 2.0)
	want := 4.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
