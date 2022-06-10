package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	args := os.Args
	file, err := os.Open(args[1])
	if err != nil {
		return err
	}
	defer file.Close()


	tmpData, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return err
	}
	data := strings.Split(string(tmpData), "\n")
	
	// sort.Strings(data)

	reverse := sort.StringSlice(data)
	sort.Sort(sort.Reverse(reverse))
	for _, str := range data {
		fmt.Println(str)
	}
	return nil
}