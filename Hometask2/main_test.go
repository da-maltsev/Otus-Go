package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testStruct struct {
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, test := range [...]testStruct{
		{
			input:    "a2d3f3",
			expected: "aadddfff",
		},
		{
			input:    "abcdefgh5j2o4",
			expected: "abcdefghhhhhjjoooo",
		},
		{
			input:    "asd123fds321a0",
			expected: "",
			err:      ErrInvalidString,
		},
	} {
		t.Run(test.input, func(t *testing.T) {
			result, err := unpack(test.input)
			require.Equal(t, test.err, err)
			require.Equal(t, test.expected, result)
		})
	}
}
