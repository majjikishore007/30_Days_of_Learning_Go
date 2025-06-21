package day3

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f and want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("%#v got %.2f want %.2f", shape, got, want)
	// 	}

	// }

	// t.Run("Rectangle", func(t *testing.T) {
	// 	shape := Shape(Rectangle{10.0, 10.0})
	// 	checkArea(t, shape, 100.0)

	// })

	// t.Run("Circle Area", func(t *testing.T) {
	// 	shape := Shape(Circle{10})
	// 	checkArea(t, shape, 214.1592653589793)

	// })
	// t.Run("Triangle Area", func(t *testing.T) {
	// 	shape := Shape(Triangle{10.0, 10.0})
	// 	checkArea(t, shape, 50.0)

	// })

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}
