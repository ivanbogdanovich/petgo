package cli

import (
	"fmt"
	"petgo/internal/geometryapp/geometry"
	"strconv"
	"strings"
)

type PointsFlag struct {
	Points []geometry.Point
}

type CircleCenterFlag struct {
	Center geometry.Point
}

func (pf *PointsFlag) String() string {
	return ""
}

func (pf *PointsFlag) Set(value string) error {
	point, err := ParsePoint(value)
	if err != nil {
		return err
	}
	pf.Points = append(pf.Points, point)
	return nil
}

func (cf *CircleCenterFlag) String() string {
	return ""
}

func (cf *CircleCenterFlag) Set(value string) error {
	point, err := ParsePoint(value)
	if err != nil {
		return err
	}
	cf.Center = point
	return nil
}

func ParsePoint(s string) (geometry.Point, error) {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return geometry.Point{}, fmt.Errorf("invalid point %q: expected x,y", s)
	}

	x, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return geometry.Point{}, fmt.Errorf("invalid X coordinate in %q: %w", s, err)
	}

	y, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return geometry.Point{}, fmt.Errorf("invalid Y coordinate in %q: %w", s, err)
	}

	return geometry.Point{X: x, Y: y}, nil
}
