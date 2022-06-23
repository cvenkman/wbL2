package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Фасад — паттерн, структурирующий объекты.
Класс Facade предоставляет унифицированный доступ для классов подсистемы
Предоставляет унифицированный интерфейс вместо набора интерфейсов
некоторой подсистемы. Фасад определяет интерфейс более высокого уровня,
который упрощает использование подсистемы.
*/

// человек может выполнять какие-то действия и о чем-то разговаривать
type Human struct {
	action Action
	talk   Talk
}

func (h *Human) DoSmth() {
	h.action.Sleep()
	h.action.Walk()
	h.talk.AboutDog()
}

/**********   классы подсистемы   **********/
type Action struct {
}

func (c *Action) Sleep() {
	fmt.Println("sleep")
}
func (c *Action) Walk() {
	fmt.Println("walk")
}

type Talk struct {
}

func (c *Talk) AboutDog() {
	fmt.Println("...dog...")
}
