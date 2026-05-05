package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStackPushPopLIFO(t *testing.T) {
	t.Parallel()

	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, ok := s.Pop()
	require.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = s.Pop()
	require.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = s.Pop()
	require.True(t, ok)
	assert.Equal(t, 1, v)

	assert.True(t, s.IsEmpty())
}
