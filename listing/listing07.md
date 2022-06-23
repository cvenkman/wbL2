Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Выведутся значения из каналов a и b, потом бесконечные нули (default int)

select {
case v, ok := <-a:
    if !ok {
        close(c)
        return
    }
    c <- v
case v, ok := <-b:
    if !ok {
        close(c)
        return
    }
    c <- v
}

ok будет истиной в случае, если канал открыт или операция чтения может быть выполнена, и горутина не заблокируется

```
