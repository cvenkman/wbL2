package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	in := []string{"Пятак", "листок", "тяпка", "листок", "пятка", "слиток", "столик", "кот"}
	m := search(in)
	fmt.Println(m)
}

// принимает массив со словами
// возврашает мапу, где:
// Ключ - первое встретившееся в словаре слово из множества
// Значение - ссылка на массив, каждый элемент которого, слово из множества
func search(dict []string) map[string][]string {
	// приводим словарь с нижнему регистру
	dict = toLowRegister(dict)

	// мапа множества анаграмм
	res := make(map[string][]string)

	for _, word := range dict {
		// флаг становится true если мы добавляем word в массив уже существующего ключа
		isAdded := false

		// проходимся по ключам и смотрим является ли какой-то ключ анаграммой word
		for key := range res {
			if isAnagram(key, word) {
				// если является - добавляем word в массив res[key]
				res[key] = append(res[key], word)
				isAdded = true
				break
			}
		}

		// если флаг false - создаем и добавляем пустой массив в новый ключ
		if isAdded == false {
			// если элемента нет
			res[word] = make([]string, 0)
		}
	}

	// проходимся по мапе и удаляем массивы с длиной 0
	for key, arr := range res {
		if len(arr) < 1 {
			delete(res, key)
		}
		// сортируем массив
		sort.Sort(sort.StringSlice(res[key]))
	}
	return res
}

// https://golangbyexample.com/check-two-strings-anagram-go/

func isAnagram(s string, t string) bool {
	// если слова одинаковые
	if s == t {
		return false
	}
	lenS := len(s)
	lenT := len(t)
	if lenS != lenT {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < lenS; i++ {
		anagramMap[string(s[i])]++
	}
	for i := 0; i < lenT; i++ {
		anagramMap[string(t[i])]--
	}

	for i := 0; i < lenS; i++ {
		if anagramMap[string(s[i])] != 0 {
			return false
		}
	}

	return true
}


func toLowRegister(arr []string) []string {
	for i := range arr {
		arr[i] = strings.ToLower(arr[i])
	}
	return arr
}