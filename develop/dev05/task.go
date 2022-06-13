package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	b bool
}

// type results struct

type result struct {
	before int
	found string
}

func main() {
	flags := flags{}
	flag.BoolVar(&flags.b, "b", false, "reverse sort")
	flag.Parse()
	args := flag.Args()
	
	toSerach := args[0]
	
	args = args[1:]
	
	// fmt.Println(args, toSerach)

	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			fmt.Fprint(os.Stderr, "grep: ", err)
			continue
		}
		data, err := getFileData(file)
		if err != nil {
			fmt.Fprint(os.Stderr, "grep: ", err)
		}
		
		// получаем массив структур с найденными словами из файла
		res := search(data, toSerach, flags)
		// fmt.Println(res)
		for _, el := range res {
			if flags.b {
				fmt.Printf("%s:%d: %s\n", file.Name(), el.before, el.found)
			} else {
				fmt.Printf("%s: %s\n", file.Name(), el.found)
			}
		}
	}

}

func search(data []string, toSerach string, flags flags) []result {
	res := make([]result, 0)
	i := 0

	// идем по массиву строк из файла
	for _, str := range data {
		if strings.Contains(str, toSerach) {
			res1 := result{before: i, found: str}
			res = append(res, res1)
		}
		// видимо включая \0
		i += len(str) + 1
	}
	return res
}

func getFileData(file *os.File) ([]string, error) {

	// if several join to one slice
	defer file.Close()
	// read file
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}

	// split file
	return strings.Split(string(data), "\n"), nil
}
