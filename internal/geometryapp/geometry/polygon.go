package geometry

import (
	"fmt"
	"math"
)

type Polygon struct {
	points []Point
}

const MinPoints = 3

func NewPolygon(points []Point) (*Polygon, error) {
	lenPoints := len(points)
	if lenPoints < MinPoints {
		return &Polygon{}, fmt.Errorf("polygon must have at least 3 points")
	}

	return &Polygon{
		points: points,
	}, nil
}

func (poly Polygon) ContainsPoint(p Point) bool {
	inside := false
	n := len(poly.points)
	j := n - 1

	for i := 0; i < n; i++ {
		xi, yi := poly.points[i].X, poly.points[i].Y
		xj, yj := poly.points[j].X, poly.points[j].Y

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
	n := len(poly.points)
	sum := 0.0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		sum += poly.points[i].X*poly.points[j].Y - poly.points[j].X*poly.points[i].Y
	}

	return math.Abs(sum) / 2
}
