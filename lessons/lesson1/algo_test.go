package lesson1

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	orig := []int{1, 2, 3, 4, 5}
	copy := slices.Clone(orig)
	res := Reverse(copy)

	assert.Equal(t, len(orig), len(res))
	slices.Reverse(res)
	assert.Equal(t, orig, res)
}

func TestDeduplicate(t *testing.T) {
	tests := map[string]struct {
		data []int
		want []int
	}{"all unique": {
		data: []int{1, 2, 3, 4, 5},
		want: []int{1, 2, 3, 4, 5},
	},

		"all same": {
			data: []int{1, 1, 1, 1},
			want: []int{1},
		},
		"has duplicates": {
			data: []int{1, 1, 2, 3},
			want: []int{1, 2, 3},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := Deduplicate(tt.data)
			assert.Equal(t, tt.want, got)
		})
	}
}
