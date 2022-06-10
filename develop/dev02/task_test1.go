package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name string
		in string
		out string
	}{
		{
			name: "a4bs2",
			in: "a4bs2",
			out: "aaaabss",
		},
		{
			name: "all chars",
			in: "abab",
			out: "abab",
		},
		{
			name: "a10b",
			in: "a10b",
			out: "a10b",
		},
		{
			name: "empty",
			in: "",
			out: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.out, Convert(testCase.in))
		})
	}
}