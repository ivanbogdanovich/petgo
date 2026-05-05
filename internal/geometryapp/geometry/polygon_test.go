package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPolygon(t *testing.T) {
	t.Parallel()

	_, err := NewPolygon([]Point{{X: 0, Y: 0}, {X: 1, Y: 1}})
	require.Error(t, err)

	poly, err := NewPolygon([]Point{
		{X: 0, Y: 0},
		{X: 2, Y: 0},
		{X: 0, Y: 2},
	})
	require.NoError(t, err)
	assert.Len(t, poly.Points, 3)
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

	assert.True(t, poly.ContainsPoint(Point{X: 2, Y: 1}))
	assert.False(t, poly.ContainsPoint(Point{X: 5, Y: 1}))
}
