package geometry

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(other Point) float64 {
	dx := other.X - p.X
	dy := other.Y - p.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p Point) IsWithinRadius(other Point, radius float64) bool {
	return p.Distance(other) <= radius
}

func ParsePoint(s string) (Point, error) {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return Point{}, fmt.Errorf("invalid point %q: expected x,y", s)
	}

	x, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return Point{}, fmt.Errorf("invalid X coordinate in %q: %w", s, err)
	}

	y, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return Point{}, fmt.Errorf("invalid Y coordinate in %q: %w", s, err)
	}

	return Point{X: x, Y: y}, nil
}