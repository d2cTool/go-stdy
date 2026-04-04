package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := NewSet[int]()
	s.Items[1] = struct{}{}
	s.Items[2] = struct{}{}
	s.Items[3] = struct{}{}
	assert.Equal(t, len(s.Items), 3)

	other := NewSet[int]()
	other.Items[3] = struct{}{}
	other.Items[4] = struct{}{}

	s.Add(other)
	assert.Equal(t, len(s.Items), 4)

	s.Remove(other)
	assert.Equal(t, len(s.Items), 2)

	s.Items[1] = struct{}{}
	s.Items[2] = struct{}{}
	s.Items[3] = struct{}{}
	res := s.Intersect(other)
	assert.Equal(t, len(res.Items), 1)
}
