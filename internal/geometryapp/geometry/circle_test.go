package geometry

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCircleContainsPoint(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		circle         Circle
		expectedResult bool
		got            bool
		point          Point
	}{
		{
			name: "TestCircleContainsPoint_Inside",
			circle: Circle{
				center: Point{X: 0, Y: 0},
				radius: 5,
			},
			point:          Point{X: 0, Y: 0},
			expectedResult: true,
		},
		{
			name: "TestCircleContainsPoint_Outside",
			circle: Circle{
				center: Point{X: 0, Y: 0},
				radius: 5,
			},
			point:          Point{X: 17, Y: 17},
			expectedResult: false,
		},
		{
			name: "TestCircleContainsPoint_ZeroRadiusDifferentPoint",
			circle: Circle{
				center: Point{X: 15, Y: 15},
				radius: 0,
			},
			point:          Point{X: 0, Y: 0},
			expectedResult: false,
		},
		{
			name: "TestCircleContainsPoint_OnBorder",
			circle: Circle{
				center: Point{X: 0, Y: 0},
				radius: 4,
			},
			point:          Point{X: 4, Y: 0},
			expectedResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.circle.ContainsPoint(tt.point)

			require.Equal(t, tt.expectedResult, got)
		})
	}
}

func TestCircleArea(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		circle         Circle
		expectedResult float64
	}{
		{
			name:           "radius 5",
			circle:         Circle{radius: 5},
			expectedResult: math.Pi * 25,
		},
		{
			name:           "radius 0.001",
			circle:         Circle{radius: 0.001},
			expectedResult: math.Pi * 0.001 * 0.001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.circle.Area()

			require.InEpsilon(t, tt.expectedResult, got, Epsilon)
		})
	}
}
