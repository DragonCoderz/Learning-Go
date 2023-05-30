package shapes

import "math"

type Shape interface {
	Area() float64
}

//Automatic interface checking

//Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface
//Circle has a method called Area that returns a float64 so it satisfies the Shape interface

type Rectangle struct {
	Width  float64
	Height float64
}

//This is a method specific to the rectangle struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

//This is a method specific to the circle struct
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * (0.5)
}

//This is a function
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
