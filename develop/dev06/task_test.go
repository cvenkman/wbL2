package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		delim          string
		fields         []int
		isWihtoutDelim bool
		res            []string
	}{
		{
			name:           "1",
			input:          "fe,fds,fd",
			delim:          ",",
			fields:         []int{1},
			isWihtoutDelim: false,
			res:            []string{"fe"},
		},
		{
			name:           "2",
			input:          "fe,fds,fd",
			delim:          ",",
			fields:         []int{1, 3},
			isWihtoutDelim: false,
			res:            []string{"fe", "fd"},
		},
		{
			name:           "3",
			input:          "fe",
			delim:          ",",
			fields:         []int{1, 3},
			isWihtoutDelim: false,
			res:            []string{"fe"},
		},
		{
			name:           "4",
			input:          "fe",
			delim:          ",",
			fields:         []int{1, 3},
			isWihtoutDelim: true,
			res:            []string(nil),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := searchResult(tc.input, tc.delim, tc.fields, tc.isWihtoutDelim)
			assert.Equal(t, tc.res, res)
		})
	}
}
