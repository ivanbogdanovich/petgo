package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueueEnqueueDequeueOrder(t *testing.T) {
	t.Parallel()

	var q Queue
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	v, ok := q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 10, v)

	v, ok = q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 20, v)

	v, ok = q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 30, v)

	require.True(t, q.IsEmpty())
	require.Equal(t, 0, q.Len())
}
