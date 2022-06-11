package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	r bool
	u bool
	c bool
}

func main() {
	flags := flags{}
	flag.BoolVar(&flags.r, "r", false, "reverse sort")
	flag.BoolVar(&flags.u, "u", false, "write only unique strings")
	flag.BoolVar(&flags.c, "c", false, "is file sorted")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("mus't be at lest 1 argument")
	}
	// если есть флаг с проверяется только первый файл
	if flags.c {
		args = args[:1]
	}
	files := openFiles(args)

	// get strings from all files
	data, err := getFilesData(files)
	if err != nil {
		log.Fatal("mus't be at lest 1 argument")
	}

	// TODO goruitne
	data = sortFile(flags, data)
	writeSTDOUT(data)

}

func sortFile(flags flags, data []string) []string {
	// just check
	if flags.c {
		// -c -r
		// просто отсортировать массив по нужным ключам и сравнить
		isSorted := sort.IsSorted(sort.StringSlice(data))
		if !isSorted {
			fmt.Println("disorder")
		}
		return []string{}
	}

	sort.Strings(data)

	if flags.r {
		sort.Sort(sort.Reverse(sort.StringSlice(data)))
	}
	if flags.u {
		data = removeDuplicateStr(data)
	}

	return data
}

func openFiles(args []string) []*os.File {
	files := make([]*os.File, 0, len(args))
	// open files
	// if several join to one slice
	for i := 0; i < len(args); i++ {
		file, err := os.Open(args[i])
		if err != nil {
			log.Fatal(err)
		}
		files = append(files, file)
	}
	return files
}

func getFilesData(files []*os.File) ([]string, error) {
	var data []string

	// if several join to one slice
	for _, file := range files {
		defer file.Close()
		// read file
		tmpData, err := ioutil.ReadFile(file.Name())
		if err != nil {
			return nil, err
		}
		// split file
		data = append(data, strings.Split(string(tmpData), "\n")...)
	}
	return data, nil
}

// write to standard output
func writeSTDOUT(data []string) {
	for _, str := range data {
		fmt.Println(str)
	}
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
