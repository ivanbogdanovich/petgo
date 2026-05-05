package geometry

import "math"

type Circle struct {
	Center Point
	Radius float64
}

func (c Circle) ContainsPoint(p Point) bool {
	return c.Center.Distance(p) <= c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
