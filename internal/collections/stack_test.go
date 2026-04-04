package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStack_IsEmpty(t *testing.T) {
	s := NewStack[int]()
	assert.True(t, s.IsEmpty())

	_, err := s.Pop()
	require.ErrorIs(t, err, ErrStackEmpty)
}

func TestStack_PushPop(t *testing.T) {
	s := NewStack[int]()
	a, b, c := 1, 2, 3
	s.Push(a)
	s.Push(b)
	s.Push(c)
	assert.False(t, s.IsEmpty())

	v, err := s.Pop()
	require.NoError(t, err)
	assert.Equal(t, 3, v)

	v, err = s.Pop()
	require.NoError(t, err)
	assert.Equal(t, 2, v)

	v, err = s.Pop()
	require.NoError(t, err)
	assert.Equal(t, 1, v)

	assert.True(t, s.IsEmpty())
}
