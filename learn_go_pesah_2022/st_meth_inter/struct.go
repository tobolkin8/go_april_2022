package st_meth_inter

import "math"

///////////////////////////----------------- Interface ------------------///////////////////////////////
type Shape interface {
	Area() float64
}

///////////////////////////----------------- Rectangle ------------------///////////////////////////////
type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

///////////////////////////----------------- Circle ------------------///////////////////////////////
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

///////////////////////////----------------- Triangle ------------------///////////////////////////////
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
