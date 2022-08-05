package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 2.0,
	}
	got := Perimeter(rectangle)
	want := 24.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func ExamplePerimeter() {
	rectangle := Rectangle{
		Width:  5.0,
		Height: 5.0,
	}
	perimeter := Perimeter(rectangle)
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

func ExampleArea() {
	area := Area(3.0, 3.0)
	fmt.Println(area)
	// Output: 9
}
