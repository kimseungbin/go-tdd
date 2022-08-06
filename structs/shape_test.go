package structs

import (
	"fmt"
	"testing"
)

type Shape interface {
	Area() float64
}

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

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{
			Width:  12,
			Height: 6,
		}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{
			Base:   12,
			Height: 6,
		}, hasArea: 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("%#v got %g want %g", test.shape, got, test.hasArea)
			}
		})
	}
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
