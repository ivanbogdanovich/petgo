package geometry

import (
	"fmt"
	"math"
)

type Polygon struct {
	Points []Point
}

func NewPolygon(points []Point) (Polygon, error) {
	if len(points) < 3 {
		return Polygon{}, fmt.Errorf("polygon must have at least 3 points")
	}
	return Polygon{Points: points}, nil
}

func (poly Polygon) ContainsPoint(p Point) bool {
	n := len(poly.Points)
	if n < 3 {
		return false
	}

	inside := false
	j := n - 1

	for i := 0; i < n; i++ {
		xi, yi := poly.Points[i].X, poly.Points[i].Y
		xj, yj := poly.Points[j].X, poly.Points[j].Y

		intersects := ((yi > p.Y) != (yj > p.Y)) &&
			(p.X < (xj-xi)*(p.Y-yi)/(yj-yi)+xi)

		if intersects {
			inside = !inside
		}

		j = i
	}

	return inside
}

func (poly Polygon) Area() float64 {
	n := len(poly.Points)
	if n < 3 {
		return 0
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		sum += poly.Points[i].X*poly.Points[j].Y - poly.Points[j].X*poly.Points[i].Y
	}

	return math.Abs(sum) / 2
}
