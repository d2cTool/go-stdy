package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNode_Options(t *testing.T) {
	data := 1
	other := NewNode[int](&data, nil, nil)
	assert.Equal(t, data, *other.data)
	assert.Nil(t, other.next)
	assert.Nil(t, other.prev)

	opt := func(n *Node[int]) {
		n.next = other
		n.prev = other
	}
	n := NewNodeWithOptions(&data, opt)

	require.NotNil(t, n.data)
	assert.Same(t, other, n.next)
	assert.Same(t, other, n.prev)
	assert.Equal(t, data, *n.data)
}
