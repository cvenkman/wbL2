package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name	string
		flags	flags
		in		[]string
		out		[]string
		err		error
	}{
		{
			name:  "simple sort (without flags)",
			flags: flags{},
			in:    []string{"c", "a", "d", "b"},
			out:   []string{"a", "b", "c", "d"},
			err:   nil,
		},
		{
			name:  "sort with blanks",
			flags: flags{},
			in:    []string{"a", " b", "c", "q"},
			out:   []string{" b", "a", "c", "q"},
			err:   nil,
		},
		//////// only one flag ////////
		{
			name:  "reverse sort (-r)",
			flags: flags{r: true},
			in:    []string{"c", "a", "d", "b"},
			out:   []string{"d", "c", "b", "a"},
			err:   nil,
		},
		{
			name:  "unique: output only the first of an equal (-u)",
			flags: flags{u: true},
			in:    []string{"c", "a", "v", "a"},
			out:   []string{"a", "c", "v"},
			err:   nil,
		},
		{
			name:  "reverse and unique sort (-r -u)",
			flags: flags{r: true, u: true},
			in:    []string{"c", "a", "d", "b", "b"},
			out:   []string{"d", "c", "b", "a"},
			err:   nil,
		},
		//////// -c and other flags ////////
		{
			name:  "check whether input is sorted (-c)",
			flags: flags{c: true},
			in:    []string{"c", "a", "d"},
			out:   nil,
			err:   errors.New("disorder"),
		},
		{
			name:  "check whether input is sorted (-c)",
			flags: flags{c: true},
			in:    []string{"a", "b", "c"},
			out:   nil,
			err:   nil,
		},
		{
			name:  "check whether input is reverse sorted (-c -r)",
			flags: flags{c: true, r: true},
			in:    []string{"a", "b", "c"},
			out:   nil,
			err:   errors.New("disorder"),
		},
		{
			name:  "check whether input is reverse sorted (-c -r)",
			flags: flags{c: true, r: true},
			in:    []string{"c", "b", "a"},
			out:   nil,
			err:   nil,
		},
		{
			name:  "check whether input is unique sorted (-c -u)",
			flags: flags{c: true, u: true},
			in:    []string{"a", "b", "c"},
			out:   nil,
			err:   nil,
		},
		{
			name:  "check whether input is unique sorted (-c -u)",
			flags: flags{c: true, u: true},
			in:    []string{"a", "b", "a", "c"},
			out:   nil,
			err:   errors.New("disorder"),
		},
		{
			name:  "check whether input is sorted (ignore leading blanks) (-c -b)",
			flags: flags{c: true, b: true},
			in:    []string{"a", "   b", "c", "d"},
			out:   nil,
			err:   nil,
		},
		//////// -b and other flags ////////
		{
			name:  "sort with ignore leading blanks (-b)",
			flags: flags{b: true},
			in:    []string{"a", "   b", "c", "a"},
			out:   []string{"a", "a", "   b", "c"},
			err:   nil,
		},
		{
			name:  "reverse sort with ignore leading blanks (-b -r)",
			flags: flags{b: true, r: true},
			in:    []string{"a", " b", "c", "d"},
			out:   []string{"d", "c", " b", "a"},
			err:   nil,
		},
		{
			name:  "ignore leading blanks and onle unisue (-b -u)",
			flags: flags{b: true, r: true},
			in:    []string{"a", " b", "c", "d"},
			out:   []string{"d", "c", " b", "a"},
			err:   nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := sortFile(testCase.flags, testCase.in)
			assert.Equal(t, testCase.out, res)
			assert.Equal(t, testCase.err, err)
		})
	}
}
