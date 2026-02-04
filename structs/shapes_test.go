package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got: %.2f want: %.2f", got, want)
	}
}

func ExamplePerimeter() {
	r := Rectangle{
		Width:  20.0,
		Height: 20.0,
	}

	res := Perimeter(r)
	fmt.Printf("%.2f", res)

	// Output: 80.00
}

func BenchmarkPerimeter(b *testing.B) {
	r := Rectangle{
		Width:  20.0,
		Height: 20.0,
	}

	for b.Loop() {
		_ = Perimeter(r)
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: &Rectangle{Width: 12.0, Height: 6.0}, want: 72.00},
		{name: "Circle", shape: &Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: &Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v got: %g want: %g", tt.shape, got, tt.want)
			}
		})
	}
}

func ExampleShape_Area() {
	r := &Rectangle{Width: 12.0, Height: 6.0}
	res := r.Area()

	fmt.Printf("%.2f", res)
	// Output: 72.00
}

func BenchmarkShape_Area(b *testing.B) {
	r := &Rectangle{Width: 12.0, Height: 6.0}

	for b.Loop() {
		_ = r.Area()
	}
}
