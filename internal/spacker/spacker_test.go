package spacker

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnpack_Table(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{
			name:  "empty",
			input: "",
			want:  "",
		},
		{
			name:  "single_letter_no_digit",
			input: "z",
			want:  "z",
		},
		{
			name:  "word_without_digits",
			input: "hello",
			want:  "hello",
		},
		{
			name:  "unicode_letters",
			input: "при2вет",
			want:  "приивет",
		},
		{
			name:  "repeat_blocks",
			input: "a2b3c",
			want:  "aabbbc",
		},
		{
			name:  "single_repeat",
			input: "x5",
			want:  "xxxxx",
		},
		{
			name:  "example1",
			input: "a4bc2d5e",
			want:  "aaaabccddddde",
		},
		{
			name:    "example2",
			input:   "aaa10b",
			wantErr: ErrInvalidInput,
		},
		{
			name:  "example3",
			input: "d\n5abc",
			want:  "d\n\n\n\n\nabc",
		},
		{
			name:    "digit_without_preceding_letter",
			input:   "2a",
			wantErr: ErrInvalidInput,
		},
		{
			name:    "digit_only",
			input:   "9",
			wantErr: ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Unpack(tt.input)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Empty(t, got)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPack_Table(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"a", "a"},
		{"aaabbc", "a3b2c"},
		{"hello", "hel2o"},
		{"при2вет", "при2вет"},
		{"a2b3", "a2b3"},
		{"aaaabccddddde", "a4bc2d5e"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, Pack(tt.input), "input %q", tt.input)
	}
}
