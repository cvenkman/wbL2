package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	args := os.Args
	if len(args) < 1 {
		log.Fatal("Usage: wget [URL]...")
	}
	urls := os.Args[1:]

	for _, url := range urls {
		_, err := Wget(url)
		if err != nil {
			log.Println(err)
		}
	}
}

// Wget is a free utility for non-interactive download of files from the Web.
// returns file with response size
func Wget(url string) (int64, error) {
	fmt.Println("Downloading ", url)

	// get response from url
	response, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return -1, fmt.Errorf("response status: %d", response.StatusCode)
	}
	fmt.Println("Response Status:", response.StatusCode)

	output := createToSave(response.Request.URL.Path)
	defer output.Close()

	// copy data from response body to file
	io.Copy(output, response.Body)
	fileinfo, err := output.Stat()
	if err != nil {
		return -1, err
	}

	fmt.Println("Size:", fileinfo.Size())
	fmt.Println("Saving to:", output.Name())
	return fileinfo.Size(), nil
}

// create file or directoey to save data
func createToSave(responsePath string) *os.File {
	if len(responsePath) > 0 {
		// delete first slash if responsePath is a directory
		if responsePath[0] == '/' {
			responsePath = responsePath[1:]
		}
	}
	// create file index.html if response.Request.URL.Path is /
	if responsePath == "" {
		responsePath = "index.html"
	}

	output, err := os.Create(responsePath)
	if err != nil {
		log.Fatal(err)
	}
	return output
}
