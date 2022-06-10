package main

import (
	// "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name string
		in string
		// out string
		err error
	}{
		{
			name: "cd -",
			in: "cd -",
			// out: "aaaabss",
			err: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := execInput(testCase.in)
			assert.Equal(t, testCase.err, err)
			_ = err
		})
	}
}