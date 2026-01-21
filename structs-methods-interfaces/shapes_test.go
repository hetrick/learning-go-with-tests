package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expected := 40.0

	if result != expected {
		t.Errorf("got %.2f, but expected %.2f", result, expected)
	}

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 10.0, Width: 10.0}, hasArea: 100.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12.0, Height: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.shape.Area()
			if result != tt.hasArea {
				t.Errorf("%#v resulted in %g, expected %g", tt.shape, result, tt.hasArea)
			}
		})
	}

	// 	checkArea := func(t testing.TB, shape Shape, expected float64) {
	// 		t.Helper()
	// 		result := shape.Area()

	// 		if result != expected {
	// 			t.Errorf("got %.2f, but expected %.2f", result, expected)
	// 		}
	// 	}
	// 	t.Run("rectangles", func(t *testing.T) {
	// 		rectangle := Rectangle{10.0, 10.0}
	// 		checkArea(t, rectangle, 100.0)

	// })
	//
	//	t.Run("circles", func(t *testing.T) {
	//		circle := Circle{10.0}
	//		checkArea(t, circle, 314.1592653589793)
	//	})
}

type Shape interface {
	Area() float64
}
