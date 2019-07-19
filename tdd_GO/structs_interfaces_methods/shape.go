package main

import "math"

func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

//Refactor the following code

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
