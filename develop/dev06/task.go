package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
 -f (field): -c option is useful for fixed-length lines. Most unix files doesn’t have
 fixed-length lines. To extract the useful information you need to cut by fields rather than columns.
 List of the fields number specified must be separated by comma. Ranges are not described with -f option.
 cut uses tab as a default field delimiter but can also work with other delimiter by using -d option.
Note: Space is not considered as delimiter in UNIX.
*/

// type data struct {
// 	inputLine string

// }

func main() {
	var field string
	var delim string
	var isWihtoutDelim bool
	flag.StringVar(&field, "f", "", "select only these fields; also print any line that contains no delimiter character, unless the -s option is specified")
	flag.StringVar(&delim, "d", "\t", "use DELIM instead of TAB for field delimiter")
	flag.BoolVar(&isWihtoutDelim, "s", false, "do not print lines not containing delimiters")
	flag.Parse()

	if field == "" {
		log.Fatal("usage: cut -f list [-s] [-d delim] [file ...]")
	}
	if delim == "" {
		delim = "\t"
	}

	writer := os.Stdin

	// какие колонки выводить
	fields := parseFiled(field)

	scanner := bufio.NewScanner(writer)
	for scanner.Scan() {
		inputLine := scanner.Text()

		res := searchResult(inputLine, delim, fields, isWihtoutDelim)
		printResult(writer, res)
	}
}

// FIXME arguments
func searchResult(inputLine, delim string, fields []int, isWihtoutDelim bool) (res []string) {
	if isWihtoutDelim && !strings.Contains(inputLine, delim) {
		return
	}
	res = make([]string, 0)
	lines := strings.Split(inputLine, delim)
	// массив со строками для результата

	for _, field := range fields {
		field--
		if field < len(lines) {
			// добавляем введеные строки в данной колонке в масив res
			res = append(res, lines[field])
		}
	}
	return
}

func printResult(writer io.Writer, result []string) {
	if len(result) == 0 {
		return

	}
	for i, str := range result {
		fmt.Fprint(writer, str)
		if i < len(result)-1 {
			fmt.Fprint(writer, ",")
		}
	}
	fmt.Fprint(writer, "\n")
}

// елси число отрицательное - добавляем все поля, если число больше кол-ва полей - игнорируем число
func parseFiled(field string) []int {
	fields := make([]int, 0)
	fieldsArr := strings.Split(field, ",")
	for _, num := range fieldsArr {
		number, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal("cut: [-f] list: illegal list value")
		} else if number == 0 {
			log.Fatal("cut: [-f] list: values may not include zero")
		}
		// не добавляем если такое число уже есть
		for _, el := range fields {
			if number == el {
				continue
			}
		}
		fields = append(fields, number)
	}
	sort.Ints(fields)
	return fields
}
