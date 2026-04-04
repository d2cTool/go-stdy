package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLruCache_InvalidCapacity(t *testing.T) {
	t.Parallel()

	for _, cap := range []int{0, -1, -100} {
		c, err := NewLruCache[string, int](cap)
		assert.Nil(t, c)
		assert.ErrorIs(t, err, ErrInvalidCapacity)
	}
}

func TestNewLruCache_OK(t *testing.T) {
	t.Parallel()

	c, err := NewLruCache[int, string](2)
	require.NoError(t, err)
	require.NotNil(t, c)
	assert.Equal(t, 2, c.cap)
}

func TestLruCache_GetMissing(t *testing.T) {
	t.Parallel()

	c, err := NewLruCache[string, int](2)
	require.NoError(t, err)

	assert.Nil(t, c.Get("nope"))
}

func TestLruCache_SetGet(t *testing.T) {
	t.Parallel()

	c, err := NewLruCache[string, int](3)
	require.NoError(t, err)

	x := 42
	c.Set("k", &x)

	got := c.Get("k")
	require.NotNil(t, got)
	assert.Equal(t, 42, *got)
}

func TestLruCache_EvictsLRU(t *testing.T) {
	t.Parallel()

	c, err := NewLruCache[string, int](2)
	require.NoError(t, err)

	a, b, cc := 1, 2, 3
	c.Set("a", &a)
	c.Set("b", &b)
	c.Set("c", &cc)

	assert.Nil(t, c.Get("a"))
	require.NotNil(t, c.Get("b"))
	require.NotNil(t, c.Get("c"))
	assert.Equal(t, 2, *c.Get("b"))
	assert.Equal(t, 3, *c.Get("c"))
}

func TestLruCache_GetPromotes(t *testing.T) {
	t.Parallel()

	c, err := NewLruCache[string, int](2)
	require.NoError(t, err)

	a, b, cc := 1, 2, 3
	c.Set("a", &a)
	c.Set("b", &b)
	require.NotNil(t, c.Get("a"))

	c.Set("c", &cc)

	assert.Nil(t, c.Get("b"))
	require.NotNil(t, c.Get("a"))
	require.NotNil(t, c.Get("c"))
}

func TestLruCache_SetUpdatesExisting(t *testing.T) {
	t.Parallel()

	c, err := NewLruCache[string, int](2)
	require.NoError(t, err)

	v1, v2 := 10, 20
	c.Set("k", &v1)
	c.Set("other", &v2)

	v1new := 99
	c.Set("k", &v1new)

	got := c.Get("k")
	require.NotNil(t, got)
	assert.Equal(t, 99, *got)
	assert.Len(t, c.items, 2)
}

func TestLruCache_Format(t *testing.T) {
	t.Parallel()

	c, err := NewLruCache[string, int](3)
	require.NoError(t, err)

	a, b := 1, 2
	c.Set("a", &a)
	c.Set("b", &b)

	s := fmt.Sprintf("%v", c)
	assert.Contains(t, s, "cap:3")
	assert.Contains(t, s, "b:")
	assert.Contains(t, s, "a:")
}
