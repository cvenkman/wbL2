Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

стек вызовов отложенных функций (с помощью defer) выполняется после оператора return
в anotherTest() значение x изменяется после того, как мы "говорим", что возвращаем x - в итоге изменяем копию x
в test() исользуется "naked" return, благодаря чему возвращается измененный x - в итоге можно изменить x

```
