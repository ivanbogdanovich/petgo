package cli

import (
	"petgo/internal/geometryapp/geometry"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsePointValid(t *testing.T) {
	t.Parallel()

	got, err := ParsePoint(" 1.5 , 2 ")
	require.NoError(t, err)
	assert.Equal(t, geometry.Point{X: 1.5, Y: 2}, got)
}

func TestParsePointInvalid(t *testing.T) {
	t.Parallel()

	_, err := ParsePoint("1,2,3")
	require.Error(t, err)
}
