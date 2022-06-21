package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name  string
		flags flags
		data  Data
		out   []string
	}{
		{
			name:  "without flags",
			flags: flags{},
			data:  Data{toSerach: "cat", strs: []string{"a a", "vfdcatvfd", "f"}, fileName: "test"},
			out:   []string{"test:vfdcatvfd"},
		},
		{
			name:  "B flag",
			flags: flags{B: true},
			data:  Data{toSerach: "cat", strs: []string{"a a", "vfdcatvfd", "", "f", "cat"}, fileName: "test"},
			out:   []string{"test:4:vfdcatvfd", "test:17:cat"},
		},
		{
			name:  "c flag",
			flags: flags{c: true},
			data:  Data{toSerach: "cat", strs: []string{"a a", "vfdcatvfd", "", "f", "cat"}, fileName: "test"},
			out:   []string{"test:2"},
		},
		{
			name:  "i & B & n",
			flags: flags{i: true, B: true, n: true},
			data:  Data{toSerach: "cat", strs: []string{"a a", "vfdcatvfd", "", "f", "cAt"}, fileName: "test"},
			out:   []string{"test:2:4:vfdcatvfd", "test:5:17:cAt"},
		},
		{
			name:  "i & C",
			flags: flags{i: true, C: true},
			data:  Data{toSerach: "cat", strs: []string{"a a", "vfdcatvfd", "", "f", "ca t", "cat"}, fileName: "test"},
			out:   []string{"test:16:vfdcatvfd", "test:22:cat"},
		},
		{
			name:  "i & n & A",
			flags: flags{i: true, n: true, A: true},
			data:  Data{toSerach: "cat", strs: []string{"a a", "vfdcatvfd", "", "f", "ca t", "cAt"}, fileName: "test"},
			out:   []string{"test:2:12:vfdcatvfd", "test:6:0:cAt"},
		},
		{
			name:  "v & n",
			flags: flags{v: true, n: true},
			data:  Data{toSerach: "cat", strs: []string{"aa", "vfdcatvfd", "", "f"}, fileName: "test"},
			out:   []string{"test:1:aa", "test:3:", "test:4:f"},
		},
		{
			name:  "F contains",
			flags: flags{F: true},
			data:  Data{toSerach: "vfdcatvfd", strs: []string{"aa", "vfdcatvfd", "", "f"}, fileName: "test"},
			out:   []string{"test:vfdcatvfd"},
		},
		{
			name:  "F",
			flags: flags{F: true},
			data:  Data{toSerach: "cat", strs: []string{"aa", "vfdcatvfd", "", "f"}, fileName: "test"},
			out:   []string{},
		},
		{
			name:  "v & F",
			flags: flags{F: true, v: true},
			data:  Data{toSerach: "vfdcatvfd", strs: []string{"aa", "vfdcatvfd", "", "f"}, fileName: "test"},
			out:   []string{"test:aa", "test:", "test:f"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			results := makeOutput(testCase.data, testCase.flags)
			assert.Equal(t, testCase.out, results)
		})
	}
}
