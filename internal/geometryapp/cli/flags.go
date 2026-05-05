package cli

import (
	"petgo/internal/geometryapp/geometry"
)

type PointsFlag struct {
	Points []geometry.Point
}

func (pf *PointsFlag) String() string {
	return ""
}

func (pf *PointsFlag) Set(value string) error {
	point, err := geometry.ParsePoint(value)
	if err != nil {
		return err
	}
	pf.Points = append(pf.Points, point)
	return nil
}

type CircleCenterFlag struct {
	Center geometry.Point
}

func (cf *CircleCenterFlag) String() string {
	return ""
}

func (cf *CircleCenterFlag) Set(value string) error {
	point, err := geometry.ParsePoint(value)
	if err != nil {
		return err
	}
	cf.Center = point
	return nil
}