package main

// import (
	// "github.com/cvenkman/wbL2/pattern"
// )

// func main() {
// 	pattern.Run()
// }








// package main

// import "fmt"

// // все утки умеют плавать;
// //суперкласс предоставляет код обобщенной реализации
// // diaplay у всех разный
// // type Duck interface {
// // 	Swim()
// // 	Display()
// // }


// type Duck struct {
// 	// ReadheadDuck
// 	// MallardheadDuck
// 	Flyable
// }

// func (d Duck) Swim() {
// 	fmt.Println("Swim: duck swim")
// }
// func (d Duck) Display() {
// 	fmt.Println("Display: duck")
// }

// type Flyable interface {
// 	Fly()
// }

// // type Quackable interface {
// // 	Quack()
// // }


// type ReadheadDuck struct {
// 	Duck
// }

// func (d ReadheadDuck) Quack() {
// 	fmt.Println("Quack: ReadheadDuck Quack")
// }
// // func (d ReadheadDuck) Swim() {
// // 	fmt.Println("Swim")
// // }
// func (d ReadheadDuck) Display() {
// 	fmt.Println("Display: ReadheadDuck")
// }
// func (d ReadheadDuck) Fly() {
// 	fmt.Println("Fly: ReadheadDuck fly")
// }

// type MallardheadDuck struct {
	
// }

// func (d MallardheadDuck) Quack() {
// 	fmt.Println("cMallardheadDuck Quack")
// }
// // func (d MallardheadDuck) Swim() {
// // 	fmt.Println("Swim")
// // }
// func (d MallardheadDuck) Display() {
// 	fmt.Println("Display: MallardheadDuck")
// }

// // func qua(q Quackable) {
// // 	q.Quack()
// // }

// func main() {
// 	// программирование на уровне реализации
// 	r := ReadheadDuck{}
// 	r.Display()
// 	r.Quack()
// 	r.Swim()
	

// 	// m := MallardheadDuck{}

// 	fmt.Println("-------")

// 	// программирование на уровне интерфейса/суперкласса
// 	d := Duck{}
// 	d.Display()
// 	d.Swim()
// 	// d.Quack()
// }