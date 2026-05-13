package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInvalidCapacity(t *testing.T) {
	t.Parallel()

	cache, err := New(0)
	require.Error(t, err)
	require.Nil(t, cache)
}

func TestCacheSetGet(t *testing.T) {
	t.Parallel()

	cache, err := New(2)
	require.NoError(t, err)

	cache.Set(1, 10)
	cache.Set(2, 20)
	cache.Set(1, 15)
	cache.Set(3, 30)

	require.Len(t, cache.Items, 2)

	_, ok := cache.Get(2)
	require.False(t, ok)

	got, ok := cache.Get(1)
	require.True(t, ok)
	require.Equal(t, 15, got)
}

func TestCacheClear(t *testing.T) {
	t.Parallel()

	cache, err := New(2)
	require.NoError(t, err)

	cache.Set(1, 10)
	cache.Set(2, 20)

	cache.Clear()

	require.Empty(t, 0, len(cache.Items))
	require.Equal(t, 0, cache.List.Size)
}
