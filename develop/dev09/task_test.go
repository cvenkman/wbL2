package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	homeURL = "http://localhost:8080"
)

func TestConvert(t *testing.T) {
	testCase := struct {
		name string
		url  string
		err  error
	}{
		name: "home page",
		url:  homeURL,
		err:  nil,
	}

	// start server with "Hello World!" page (size 12)
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World!")
		})
		http.ListenAndServe(":8080", nil)
	}()

	t.Run(testCase.name, func(t *testing.T) {
		size, err := Wget(testCase.url)
		fmt.Println(testCase.url)
		assert.Equal(t, int64(len("Hello World!")), size)
		assert.Equal(t, testCase.err, err)
	})

}
