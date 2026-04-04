package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_Options(t *testing.T) {
	data := 1
	other := NewNode(data)
	assert.Equal(t, data, other.data)
	assert.Nil(t, other.next)
	assert.Nil(t, other.prev)
}
