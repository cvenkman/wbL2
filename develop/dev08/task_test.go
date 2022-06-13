package main

import (
	// "errors"
	"errors"
	// "log" 
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	userEnv := os.Getenv("USER")

	testCases := []struct {
		name string
		in []string
		out []string
		err error
	}{
		{
			name: "echo",
			in: []string{"echo", "hi"},
			out: []string{"echo", "hi"},
			err: nil,
		},
		{
			name: "echo with \"\" ",
			in: []string{"ECHO", "\"h\"", "i"},
			out: []string{"ECHO", "h", "i"},
			err: nil,
		},
		{
			name: "echo with \" ",
			in: []string{"echo", "\"h", "i"},
			out: nil,
			err: errors.New("unclosed quote"),
		},
		{
			name: "echo with env",
			in: []string{"echo", "$USER"},
			out: []string{"echo", userEnv},
			err: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := echo(testCase.in)
			assert.Equal(t, testCase.out, actual)
			assert.Equal(t, testCase.err, err)
		})
	}
}


func TestCD(t *testing.T) {
	home := os.Getenv("HOME")

	testCases := []struct {
		name string
		in []string
		out string
		err error
	}{
		{
			name: "cd",
			in: []string{"cd"},
			out: home,
			err: nil,
		},
	}

	// file, err := os.Create("output")
	// if err != nil {
	// 	log.Fatal("can't create file for test output:", err)
	// }
	// os.Stdout = file
	// execInput("pwd")
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := cd(testCase.in)
			assert.Equal(t, testCase.out, actual)
			assert.Equal(t, testCase.err, err)
		})
	}
}


// func TestAll(t *testing.T) {
// 	file, err := os.Create("output")
// 	if err != nil {
// 		log.Fatal("can't create file for test output:", err)
// 	}
// 	defer file.Close()
// 	os.Stdout = file
// 	os.Stderr = file

// 	run("ls")
// }