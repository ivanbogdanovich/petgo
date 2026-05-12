package cli

import (
	"flag"
	"fmt"
	"petgo/internal/geometryapp/geometry"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(args []string) error {
	fs := flag.NewFlagSet("geomcli", flag.ContinueOnError)

	var pointsFlag PointsFlag
	var isCircle bool
	var centerCircle CircleCenterFlag
	var radiusCircle float64
	var checkInside bool

	fs.Var(&pointsFlag, "points", "polygon points or a point to check, format: x,y")
	fs.BoolVar(&isCircle, "circle", false, "use circle mode")
	fs.Var(&centerCircle, "center", "circle center, format: x,y")
	fs.Float64Var(&radiusCircle, "radius", 0.0, "circle radius")
	fs.BoolVar(&checkInside, "check-inside", false, "check whether point is inside shape")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	if isCircle {
		circle, err := geometry.NewCircle(centerCircle.Center, radiusCircle)
		if err != nil {
			return fmt.Errorf("failed to create circle: %w", err)
		}

		fmt.Printf("Circle area: %.2f\n", circle.Area())

		if checkInside {
			if len(pointsFlag.Points) == 0 {
				return fmt.Errorf("point is required for --check-inside in circle mode")
			}

			p := pointsFlag.Points[0]
			fmt.Printf("Point %+v inside circle: %v\n", p, circle.ContainsPoint(p))
		}

		return nil
	}

	polygon, err := geometry.NewPolygon(pointsFlag.Points)
	if err != nil {
		return fmt.Errorf("failed to create polygon: %w", err)
	}

	fmt.Printf("Polygon area: %.2f\n", polygon.Area())

	if checkInside {
		return fmt.Errorf("--check-inside only for circle")
	}

	return nil
}
