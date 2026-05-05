package geometry

import (
	"math"
	"testing"
)

func TestCircleContainsPoint(t *testing.T) {
    circle1 := Circle{
                Center: Point{X: 0, Y: 0},
                Radius: 5,
            }
    circle2 := Circle{
                Center: Point{X: 0, Y: 0},
                Radius: 5,
            }
    circle3 := Circle{
                Center: Point{X: 15, Y: 15},
                Radius: 0,
            }
    circle4 := Circle{
                Center: Point{X: 0, Y: 0},
                Radius: 4,
            }
    tests := []struct{
        name string
        circle Circle
        want bool
        got bool
    }{
        {
            name: "TestCircleContainsPoint_Inside",
            circle: circle1,
            got:  circle1.ContainsPoint(Point{X: 0, Y: 0}),
            want: true,
        },
        {
            name: "TestCircleContainsPoint_Outside",
            circle: circle2,
            got: circle2.ContainsPoint(Point{X: 17, Y: 17}),
            want: false,
        },
        {
            name: "TestCircleContainsPoint_ZeroRadiusDifferentPoint",
            circle: circle3,
            got: circle3.ContainsPoint(Point{X: 0, Y: 0}),
            want: false,
        },
        {
            name: "TestCircleContainsPoint_OnBorder",
            circle: circle4,
            got: circle4.ContainsPoint(Point{X: 4, Y: 0}),
            want: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.got != tt.want {
		        t.Fatalf("got %v, want %v", tt.got, tt.want)
	        }
        })
    }
}

const eps = 1e-9

func almostEqual(a, b float64) bool {
    return math.Abs(a - b) < eps
}

func TestCircleArea(t *testing.T) {
    tests := []struct{
        name string
        circle Circle
        want float64
    }{
        {
            name: "radius 0",
            circle: Circle{Radius: 0},
            want: 0,
        },
        {
            name: "radius 0.001",
            circle: Circle{Radius: 0.001},
            want: math.Pi * 0.001 * 0.001,
        },
    }

    for _, tt := range tests {
        got := tt.circle.Area()
        t.Run(tt.name, func(t *testing.T) {
            if(!almostEqual(got, tt.want)) {
                t.Fatalf("got %v, want %v", got, tt.want)
            }
        })
    }
}