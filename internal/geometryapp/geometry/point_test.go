package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPointIsWithinRadius(t *testing.T) {
	t.Parallel()

	center := Point{X: 0, Y: 0}
	inside := Point{X: 3, Y: 4}
	outside := Point{X: 6, Y: 0}

	assert.True(t, center.IsWithinRadius(inside, 5))
	assert.False(t, center.IsWithinRadius(outside, 5))
}

func TestParsePointValid(t *testing.T) {
	t.Parallel()

	got, err := ParsePoint(" 1.5 , 2 ")
	require.NoError(t, err)
	assert.Equal(t, Point{X: 1.5, Y: 2}, got)
}

func TestParsePointInvalid(t *testing.T) {
	t.Parallel()

	_, err := ParsePoint("1,2,3")
	require.Error(t, err)
}
