package geometry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPolygon(t *testing.T) {
	t.Parallel()

	poly, err := NewPolygon([]Point{
		{X: 0, Y: 0},
		{X: 2, Y: 0},
		{X: 0, Y: 2},
	})
	require.NoError(t, err)
	require.Len(t, poly.points, 3)
}

func TestPolygonContainsPoint(t *testing.T) {
	t.Parallel()

	poly, err := NewPolygon([]Point{
		{X: 0, Y: 0},
		{X: 4, Y: 0},
		{X: 4, Y: 3},
		{X: 0, Y: 3},
	})
	require.NoError(t, err)

	require.True(t, poly.ContainsPoint(Point{X: 2, Y: 1}))
	require.False(t, poly.ContainsPoint(Point{X: 5, Y: 1}))
}
