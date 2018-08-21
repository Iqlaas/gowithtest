package main

import (
	"math"
)

func main() {}

// Perimeter function
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

// Area function
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.width * rectangle.height
// }

// Rectangle type struct
type Rectangle struct {
	width  float64
	height float64
	// in tutorial, its capital. but it works with small.
	// afaik on time of writing.
}

// Area method for rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Circle type struct
type Circle struct {
	Radius float64
}

// Area method for circle
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Shape interface
type Shape interface {
	Area() float64
}
