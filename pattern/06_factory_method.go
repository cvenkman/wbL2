package pattern

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Фабричный метод — паттерн, порождающий классы.

	Определяет интерфейс для создания объекта, но оставляет подклассам решение о том, экземпляры
	какого класса должны создаваться. Фабричный метод позволяет классу делегировать создание экземпляров подклассам.
*/

type file string

const (
	json file = "json"
	xml  file = "xml"
)

/********   product   *********/

// интерфейс продукта
type Product interface {
	// каждый файл должен уметь возвращать сове расширение
	GetExtension() string
}

// конкретная реализация продукта
type JsonFile struct {
}

func (f *JsonFile) GetExtension() string {
	return "json"
}

// конкретная реализация продукта
type XmlFile struct {
}

func (f *XmlFile) GetExtension() string {
	return "xml"
}

/********   Creator   *********/

// Creator provides a factory interface.
type Creator interface {
	FactoryMethod(f file) Product // create file
}

// ConcreteCreator implements Creator interface.
type ConcreteCreator struct{}

// NewCreator is the ConcreteCreator constructor.
func NewCreator() Creator {
	return &ConcreteCreator{}
}

// CreateProduct is a Factory Method.
func (p *ConcreteCreator) FactoryMethod(f file) Product {
	var product Product

	switch f {
	case json:
		product = &JsonFile{}
	case xml:
		product = &XmlFile{}
	default:
		log.Fatalln("Unknown file")
	}

	return product
}

func UsageFactory() {
	factory := NewCreator()
	files := []Product{
		factory.FactoryMethod(json),
		factory.FactoryMethod(xml),
		factory.FactoryMethod(xml),
	}
	for _, file := range files {
		fmt.Println(file.GetExtension())
	}
}
