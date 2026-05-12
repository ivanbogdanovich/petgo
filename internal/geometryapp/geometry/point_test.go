package geometry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPointDistance(t *testing.T) {
	t.Parallel()

	p1 := Point{X: 0, Y: 0}
	p2 := Point{X: 3, Y: 4}

	require.InEpsilon(t, 5.0, p1.Distance(p2), Epsilon)
	require.InEpsilon(t, 5.0, p2.Distance(p1), Epsilon)
}
