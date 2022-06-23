package pattern

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Паттерн Command относится к поведенческим паттернам уровня объекта.
Паттерн Command позволяет представить запрос в виде объекта. Из этого следует, что команда - это объект.
*/

//Базовый абстрактный класс Command описывающий интерфейс команды
type Command interface {
	GetTime()
}

//Класс ConcreteCommand, реализующий команду
type Clock struct {
	currentTime time.Time
}

func (c *Clock) GetTime() {
	fmt.Println("previous:", c.currentTime)
	c.currentTime = time.Now()
	fmt.Println("now:", c.currentTime)
}

// Класс Invoker, реализующий инициатора, записывающий команду и провоцирующий её выполнение
type CommandSender struct {
	command command
}

func (c *CommandSender) Send() {
	c.command.execute()
}

type command interface {
	execute()
}

type SetTimeCommand struct {
	target target
}

func (c *SetTimeCommand) execute() {
	c.target.NowTime()
}

type GetTimeCommand struct {
	target target
}

func (c *GetTimeCommand) execute() {
	c.target.GetTime()
}

type target interface {
	NowTime()
	GetTime()
}
