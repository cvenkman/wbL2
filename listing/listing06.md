Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
[3 2 3]

Первый элемент слайса изменился потому что в modifySlice мы передаем слайс s по знаечнию, но т.к. внутри слайса ссылка на массив (там три параметра: len, cap, ссылка на массив), в modifySlice можно изменить и сам слайс (но не len и cap).

Длина и второй элемент не изменились т.к. после первого append i - новый слайс и его ссылка на массив - новая ссылка

```
