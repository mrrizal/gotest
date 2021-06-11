package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("test perimeter", func(t *testing.T) {
		rectangle := Rectangle{Height: 10, Width: 10}

		got := Perimeter(rectangle)
		want := 40.0

		assert.Equal(t, got, want)
	})

	t.Run("test perimeter", func(t *testing.T) {
		rectangle := Rectangle{Height: 12.0, Width: 3.0}
		got := Perimeter(rectangle)
		want := 30.0

		assert.Equal(t, got, want)
	})
}

func TestArea(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			assert.Equal(t, got, tt.hasArea)
		})
	}
}
