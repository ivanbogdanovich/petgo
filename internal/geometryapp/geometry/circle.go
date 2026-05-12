package geometry

import (
	"fmt"
	"math"
)

type Circle struct {
	center Point
	radius float64
}

func NewCircle(center Point, radius float64) (*Circle, error) {
	if radius < 0 {
		return nil, fmt.Errorf("radius must be positive")
	}

	return &Circle{
		center: center,
		radius: radius,
	}, nil
}

func (c Circle) ContainsPoint(p Point) bool {
	return c.center.Distance(p) <= c.radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
