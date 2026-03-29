package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := NewSet[int]()
	s.items[1] = struct{}{}
	s.items[2] = struct{}{}
	s.items[3] = struct{}{}
	assert.Equal(t, len(s.items), 3)

	other := NewSet[int]()
	other.items[3] = struct{}{}
	other.items[4] = struct{}{}

	s.Add(other)
	assert.Equal(t, len(s.items), 4)

	s.Remove(other)
	assert.Equal(t, len(s.items), 2)

	s.items[1] = struct{}{}
	s.items[2] = struct{}{}
	s.items[3] = struct{}{}
	res := s.Intersect(other)
	assert.Equal(t, len(res.items), 1)
}
