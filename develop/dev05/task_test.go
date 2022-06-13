package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name	string
		flags	flags
		in		[]string
		out		[]result
	}{
		// {
		// 	name:  "grep one file",
		// 	flags: flags{},
		// 	in:    []string{"aa", "cat", "\n","d", "b"},
		// 	out:   []string{"\n", "a", "b", "c", "d"},
		// 	err:   nil,
		// },
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// res, err := sortFile(testCase.flags, testCase.in)
			assert.Equal(t, testCase.out, res)
			assert.Equal(t, testCase.err, err)
		})
	}
}
