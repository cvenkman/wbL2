package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"strconv"
	"unicode"
)

func main() {
	// fmt.Println(Convert("a4bs2"))
	// fmt.Println(Convert("a10b"))
	// fmt.Println(Convert("abab"))
	// fmt.Println(Convert("45"))
	// fmt.Println(Convert(""))
}

func Convert(before string) (string, error) {
	if len(before) == 0 {
		return before, nil
	}

	res := make([]rune, 0, len(before))

	_, err := strconv.Atoi(string(before[0]))
	if err == nil {
		return "", errors.New("must be at least one char")
	}

	i := 0

	for ; i < len(before) - 1; i++ {
		// if next char is not a digit
		if !unicode.IsDigit(rune(before[i + 1])) {
			res = append(res, rune(before[i]))
			continue
		}
		
		num, numCount := parseNumber(before[i+1:])
		for ; num > 0; num-- {
			res = append(res, rune(before[i]))
		}
		i += numCount
	}

	// add last char if it's not a digit
	if i < len(before) {
		res = append(res, rune(before[i]))
	}
	
	return string(res), nil
}

// возвращает число и количсевто цифр в нем
func parseNumber(str string) (int, int) {
	num := 0

	for i, ch := range str {
		tmpNum, err := strconv.Atoi(string(ch))
		if err != nil {
			return num, i
		}
		num = num * 10 + tmpNum
	}
	return num, len(str) + 1
}