package doubly_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListScenario(t *testing.T) {
	t.Parallel()

	var list List

	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)
	assert.Equal(t, []int{10, 20, 30}, valuesForward(list))
	assert.Equal(t, []int{30, 20, 10}, valuesBackward(list))

	list.PushFront(5)
	list.PushFront(1)
	assert.Equal(t, []int{1, 5, 10, 20, 30}, valuesForward(list))
	assert.Equal(t, []int{30, 20, 10, 5, 1}, valuesBackward(list))

	v, ok := list.PopFront()
	require.True(t, ok)
	assert.Equal(t, 1, v)
	assert.Equal(t, []int{5, 10, 20, 30}, valuesForward(list))

	v, ok = list.PopBack()
	require.True(t, ok)
	assert.Equal(t, 30, v)
	assert.Equal(t, []int{5, 10, 20}, valuesForward(list))
}

func valuesForward(l List) []int {
	var values []int
	for cur := l.Head; cur != nil; cur = cur.Next {
		values = append(values, cur.Value)
	}
	return values
}

func valuesBackward(l List) []int {
	var values []int
	for cur := l.Tail; cur != nil; cur = cur.Prev {
		values = append(values, cur.Value)
	}
	return values
}
