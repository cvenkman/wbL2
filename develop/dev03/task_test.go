package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name  string
		flags flags
		in    []string
		out   []string
	}{
		{
			name:  "reverse",
			flags: flags{r: true},
			in:    []string{"c", "a", "d", "b"},
			out:   []string{"d", "c", "b", "a"},
		},
		{
			name:  "simple",
			flags: flags{},
			in:    []string{"c", "a", "d", "b"},
			out:   []string{"a", "b", "c", "d"},
		},
		{
			name:  "unique",
			flags: flags{u: true},
			in:    []string{"c", "a", "v", "a"},
			out:   []string{"a", "c", "v"},
		},
		{
			name:  "reverse and unique",
			flags: flags{r: true, u: true},
			in:    []string{"c", "a", "d", "b", "b"},
			out:   []string{"d", "c", "b", "a"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := sortFile(testCase.flags, testCase.in)
			assert.Equal(t, testCase.out, res)
		})
	}
}
