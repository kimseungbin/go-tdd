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

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  12,
			Height: 6,
		}
		got := Area(rectangle)
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{Radius: 10}
		got := Area(circle)
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func ExampleArea() {
	rectangle := Rectangle{
		Width:  3.0,
		Height: 3.0,
	}
	area := Area(rectangle)
	fmt.Println(area)
	// Output: 9
}
