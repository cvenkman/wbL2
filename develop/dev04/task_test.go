package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name string
		in []string
		out map[string][]string
	}{
		{
			name: "1",
			in: []string{"Пятак", "листок", "тяпка", "листок", "пятка", "слиток", "столик", "кот"},
			out: map[string][]string{"листок": {"слиток", "столик"}, "пятак": {"пятка", "тяпка"}},
		},
		{
			name: "2",
			in: []string{"Пятак"},
			out: map[string][]string{},
		},
		{
			name: "3",
			in: []string{"тяПкА", "ПяТак"},
			out: map[string][]string{"тяпка": {"пятак"}},
		},
		{
			name: "4",
			in: []string{},
			out: map[string][]string{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := search(testCase.in)
			assert.Equal(t, testCase.out, res)
		})
	}
}