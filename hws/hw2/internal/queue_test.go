package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueue_IsEmpty(t *testing.T) {
	q := NewQueue[int]()
	assert.True(t, q.IsEmpty())

	_, err := q.Dequeue()
	require.ErrorIs(t, err, ErrQueueEmpty)
}

func TestQueue_AddRemove(t *testing.T) {
	q := NewQueue[int]()
	a, b, c := 1, 2, 3

	q.Enqueue(&a)
	assert.False(t, q.IsEmpty())

	q.Enqueue(&b)
	q.Enqueue(&c)

	d, err := q.Dequeue()
	require.NoError(t, err)
	require.NotNil(t, d)
	assert.Equal(t, 1, *d)

	d, err = q.Dequeue()
	require.NoError(t, err)
	require.NotNil(t, d)
	assert.Equal(t, 2, *d)

	d, err = q.Dequeue()
	require.NoError(t, err)
	require.NotNil(t, d)
	assert.Equal(t, 3, *d)

	assert.True(t, q.IsEmpty())
}
