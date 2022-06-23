package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Паттерн Strategy относится к поведенческим паттернам уровня объекта.
Паттерн Strategy определяет набор алгоритмов схожих по роду деятельности, инкапсулирует
их в отдельный класс и делает их подменяемыми. Паттерн Strategy позволяет подменять алгоритмы без участия клиентов, которые используют эти алгоритмы.
*/

// Абстрактный класс Strategy, определяющий интерфейс различных стратегий
type StrategyPrint interface {
	Print(string)
}

// печатает строку в прямом порядке
type JustPrint struct {
}

func (p *JustPrint) Print(str string) {
	fmt.Println(str)
}

// печатает строку с пробелом после каждого символа
type PrintWithSpace struct {
}

func (p *PrintWithSpace) Print(str string) {
	for _, char := range str {
		fmt.Print(char)
		fmt.Println(" ")
	}
}

// Класс Context, представляющий собой контекст выполнения той или иной стратегии
type Context struct {
	strategy StrategyPrint
}

func (c *Context) Print(str string) {
	c.strategy.Print(str)
}
