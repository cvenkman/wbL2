Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil

	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

type error interface {
    Error() string
}

Foo() возвращает [nil, *os.PathError] - значение, динамический тип. Сравниваем с [nil, nil]

Интерфейс внутри:
type iface struct {
    tab  *itab
    data unsafe.Pointer
}

Объект интерфейса в Go содержит два поля: tab с информацией о конкретном типе и data, где лежит ссылка на сами данные. И вот, по правилам Go, интерфейс может быть равен nil только если оба этих поля не определены.

Вывод: var err внутри имеет тип *os.PathError, поэтому не может равняться nil

Использование встроенного типа интерфейса error (var err error = nil) вместо конкретного типа (var err *os.PathError) исправит ситуацию

```

https://habr.com/ru/post/597461/

https://habr.com/ru/post/449714/