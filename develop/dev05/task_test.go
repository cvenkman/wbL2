package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name     string
		flags    flags
		data     Data
		out      []string
	}{
		{
			name:     "one",
			flags:    flags{},
			data:     Data{toSerach: "cat", strs: []string{"a a", "vfdcatvfd", "f"}, fileName: "test"},
			out: []string{"test: vfdcatvfd"},
		},
		{
			name:     "B flag",
			flags:    flags{B: true},
			data:     Data{toSerach: "cat", strs: []string{"a a", "vfdcatvfd", "", "f", "cat"}, fileName: "test"},
			out: []string{"test:4: vfdcatvfd", "test:17: cat"},
		},
		{
			name:     "c flag",
			flags:    flags{c: true},
			data:     Data{toSerach: "cat", strs: []string{"a a", "vfdcatvfd", "", "f", "cat"}, fileName: "test"},
			out: []string{"test: 2"},
		},
		///// n & i /////
		{
			name:     "i flag",
			flags:    flags{i: true, B: true},
			data:     Data{toSerach: "cat", strs: []string{"a a", "vfdcatvfd", "", "f", "cAt"}, fileName: "test"},
			out: []string{"test:4: vfdcatvfd", "test:17: cAt"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			results := makeOutput(testCase.data, testCase.flags)
			assert.Equal(t, testCase.out, results)
		})
	}
}
