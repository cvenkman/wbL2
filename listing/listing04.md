Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
.
.
9
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/cvenkman/Desktop/wbL2/main.go:11 +0xa8
exit status 2

for печатает все значения из канала, а потом блокируется
чтобы исправить, нужно после цикла в горутине закрыть канал (close(ch)),
тогда for range по каналу будет до тех пор, пока канал не закроется
```
