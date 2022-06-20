package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
	A bool
	B bool
	C bool
	c bool
	i bool
	v bool
	n bool
}

// type results struct

type result struct {
	lenBefore  int
	lenAfter   int
	lenContext int
	strNumber  int
	found      string
}

type Data struct {
	fileName string
	strs     []string
	toSerach string
}

func main() {
	flags := flags{}
	flag.BoolVar(&flags.B, "B", false, "reverse sort")
	flag.BoolVar(&flags.A, "A", false, "reverse sort")
	flag.BoolVar(&flags.C, "C", false, "reverse sort")
	flag.BoolVar(&flags.c, "c", false, "Prints only a count of the lines that contain the pattern.")
	flag.BoolVar(&flags.i, "i", false, "Ignores upper/lower case distinction during comparisons.")
	flag.BoolVar(&flags.n, "n", false, "Precedes each line by its line number in the file.")
	flag.BoolVar(&flags.v, "v", false, "Prints all lines except those that contain the pattern")
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("err args: grep [OPTIONS] PATTERN [FILE...]")
	}

	args := flag.Args()
	toSerach := args[0]
	args = args[1:]

	// цикл по файлам
	for _, fileName := range args {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprint(os.Stderr, "grep: ", err)
			continue
		}
		fileData := Data{}
		fileData.strs, err = getFileData(file)
		if err != nil {
			fmt.Fprint(os.Stderr, "grep: ", err)
			continue
		}
		fileData.fileName = file.Name()
		fileData.toSerach = toSerach

		results := makeOutput(fileData, flags)
		for _, res := range results {
			fmt.Println(res)
		}
	}

}

// возвращает массив со строками на вывод
func makeOutput(data Data, flags flags) []string {
	// получаем массив структур с найденными словами из файла
	results := search(data, flags)

	// структура с выводом
	out := make([]string, 0)

	// сохраняем только количество строк
	if flags.c {
		out = append(out, fmt.Sprintf("%s:%d", data.fileName, len(results)))
		return out
	}

	// сохраняем все результаты из одного файла
	for _, result := range results {
		str := data.fileName + ":"

		if flags.n {
			str += fmt.Sprintf("%d:", result.strNumber)
			// out = append(out, fmt.Sprintf("%s:%d: %s", data.fileName, result.strNumber, result.found))
		}
		if flags.B {
			// out = append(out, fmt.Sprintf("%s:%d: %s", data.fileName, result.lenBefore, result.found))
			str += fmt.Sprintf("%d:", result.lenBefore)
		}
		if flags.A {
			str += fmt.Sprintf("%d:", result.lenAfter)
			// out = append(out, fmt.Sprintf("%s:%d: %s", data.fileName, result.lenAfter, result.found))
		}
		if flags.C {
			str += fmt.Sprintf("%d:", result.lenBefore + result.lenAfter)
			// out = append(out, fmt.Sprintf("%s:%d: %s", data.fileName, result.lenAfter+result.lenBefore, result.found))
		}

		str += result.found
		out = append(out, str)
	}
	return out
}

// возвращает массив со сторокой где найдено слова и значниями для флага
func search(data Data, flags flags) []result {
	res := make([]result, 0)
	i := 0

	// идем по массиву строк из файла
	for j, str := range data.strs {
		if flags.i {
			str = strings.ToLower(str)
		}
		if (strings.Contains(str, data.toSerach) && !flags.v) ||
			(flags.v && !strings.Contains(str, data.toSerach)) {

			l := getFileDataLen(data.strs) - i - len(data.strs[j])
			res1 := result{lenBefore: i, found: data.strs[j], lenAfter: l, strNumber: j + 1}
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

func getFileDataLen(data []string) int {
	l := 0
	for _, str := range data {
		l += len(str) + 1
	}
	return l - 1
}
