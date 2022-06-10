package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name string
		in string
		out string
		err error
	}{
		{
			name: "letters and numbers",
			in: "a4bs2",
			out: "aaaabss",
			err: nil,
		},
		{
			name: "all letters",
			in: "abab",
			out: "abab",
			err: nil,
		},
		{
			name: "letter at the end",
			in: "a10b",
			out: "aaaaaaaaaab",
			err: nil,
		},
		{
			name: "empty string",
			in: "",
			out: "",
			err: nil,
		},
		{
			name: "one letter",
			in: "a",
			out: "a",
			err: nil,
		},
		{
			name: "error: all numbers",
			in: "45",
			out: "",
			err: errors.New("must be at least one char"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := Convert(testCase.in)
			assert.Equal(t, testCase.out, res)
			assert.Equal(t, testCase.err, err)
			_ = err
		})
	}
}