package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Паттерн Visitor относится к поведенческим паттернам уровня объекта.

Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами, а также позволяет добавить новый метод в класс объекта,
при этом, не изменяя сам класс этого объекта.
*/

//Абстрактный класс Visitor, описывающий интерфейс визитера
type Visitor interface {
	VisitRectangle(rectangle *Rectangle)
	VisitCircle(circle *Circle)
}

//Класс ConcreteVisitor, реализующий конкретного визитера.
//Реализует методы для обхода конкретного элемента
type ConcreteVisitor struct{}

func (ConcreteVisitor) VisitRectangle(rectangle *Rectangle) {
	fmt.Printf("Visitor visited Rectangle: %+v\n", *rectangle)
}

func (ConcreteVisitor) VisitCircle(circle *Circle) {
	fmt.Printf("Visitor visited Circle: %+v\n", *circle)
}

type Shapes struct {
	rectangle Rectangle
	circle    Circle
}

type Rectangle struct {
	width, height int
}

// Circle объект круга
type Circle struct {
	radius int
}

// Принимает интерфейс посетителя
func (s *Shapes) Visit(v Visitor) {
	v.VisitRectangle(&s.rectangle)
	v.VisitCircle(&s.circle)
}
