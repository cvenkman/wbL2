package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
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
	b bool
}

func main() {
	flags := flags{}
	flag.BoolVar(&flags.r, "r", false, "reverse sort")
	flag.BoolVar(&flags.u, "u", false, "write only unique strings")
	flag.BoolVar(&flags.c, "c", false, "is file sorted")
	flag.BoolVar(&flags.b, "b", false, "ignore blanks")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("mus't be at least 1 argument")
	}
	// если есть флаг -с проверяется только первый файл, остальные игнорируются
	if flags.c {
		args = args[:1]
	}
	files := openFiles(args)

	// get strings from all files
	data, err := getFilesData(files)
	if err != nil {
		log.Fatal("mus't be at least 1 argument")
	}

	// TODO goruitne
	data, err = sortFile(flags, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	writeSTDOUT(data)
}

// return error if flag c and array not sorted
func sortFile(flags flags, data []string) ([]string, error) {

	// just check is file sorted, not sort
	if flags.c {
		return nil, flagC(flags, data)
	}

	return sortWithFlags(flags, data), nil
}

// return error if data is not sorted
func flagC(flags flags, data []string) error {
	// just check
	// -c -r

	// cretate new slice with data
	sortedData := make([]string, len(data))
	copy(sortedData, data)
	// sort new slice with flags
	sortedData = sortWithFlags(flags, sortedData)

	// compare sorted slice with data slice
	// if false - data is not sorted with flags
	if !reflect.DeepEqual(data, sortedData) {
		return errors.New("disorder")
	}

	return nil
}

func sortWithFlags(flags flags, dataToSort []string) []string {
	if !flags.b && !flags.r {
		sort.Strings(dataToSort)
	}

	// ignore blanks
	if flags.b {
		sort.Slice(dataToSort, func(i, j int) bool {
			dataI := strings.TrimSpace(dataToSort[i])
			dataJ := strings.TrimSpace(dataToSort[j])
			return dataI < dataJ
		})
		// return dataToSort
	}
	
	// reverse sort
	if flags.r {
		sort.Slice(dataToSort, func(i, j int) bool {
			if flags.b {
				dataI := strings.TrimSpace(dataToSort[i])
				dataJ := strings.TrimSpace(dataToSort[j])
				return !(dataI < dataJ)
			}
			return !(dataToSort[i] < dataToSort[j])
		})
		// sort.Sort(sort.Reverse(sort.StringSlice(dataToSort)))
	}


	/// TODO -u -b
	// only unique strings
	if flags.u {
		dataToSort = removeDuplicateStr(dataToSort, flags.b)

	}
	return dataToSort
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

func removeDuplicateStr(strSlice []string, ignoreBlanks bool) []string {
	allKeys := make(map[string]bool)
	list := []string{}

	// for _, item := range strSlice {
	for i := 0; i < len(strSlice); i++ {
		item := strSlice[i]
		
		// если такого элемента нет в мапе - добавляем
		// если флаг -b то должны еще проверить с пробелами
		if _, value := allKeys[item]; !value {
			isAppend := true

			if ignoreBlanks {
				itemWithoutSpaces := strings.TrimSpace(item)
				for key := range allKeys {
					if strings.TrimSpace(key) == itemWithoutSpaces {
						isAppend = false
					}
				}
			}

			if isAppend {
				allKeys[item] = true
				list = append(list, item)
			}
		}
	}
	return list
}

// write to standard output
func writeSTDOUT(data []string) {
	for _, str := range data {
		fmt.Println(str)
	}
}