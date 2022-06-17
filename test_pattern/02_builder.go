package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Паттерн Строитель предлагает вынести конструирование объекта за пределы его собственного класса,
поручив это дело отдельным объектам, называемым строителями.
*/

// что должен уметь делать строитель
// объявляет шаги конструирования продуктов, общие для всех видов строителей.
type Builder interface {
	SetWalls()
	SetRoof()
	SetPool()
}


type Director struct {
	// builder Builder
}

func (d *Director) NewWoodenHouse(builder *WoodenHouseBuilder) {
	builder.SetWalls()
	builder.SetRoof()
	builder.SetPool()
}

// нет бассейна
func (d *Director) NewStoneHouse(builder *StoneHouseBuilder) {
	builder.SetWalls()
	builder.SetRoof()
}

func NewDirector() *Director {
	return &Director{}
}



// Конкретные строители реализуют строительные шаги, каждый по-своему.
// Конкретные строители могут производить разнородные объекты, не имеющие общего интерфейса.

////////   1   /////////

// конкертная реализация строителя
type StoneHouseBuilder struct {
	walls string
	roof string
}

func (b *StoneHouseBuilder) SetWalls() {
	b.roof = "roof"
}

func (b *StoneHouseBuilder) SetRoof() {
	b.walls = "wooden walls"
}

func (b *StoneHouseBuilder) GetHouse() {
	return
}

////////   2   /////////
type WoodenHouseBuilder struct {
	walls string
	roof string
	pool string
}

func (b *WoodenHouseBuilder) SetWalls() {
	b.walls = "walls"
}

func (b *WoodenHouseBuilder) SetRoof() {
	b.roof = "wooden roof"
}

func (b *WoodenHouseBuilder) SetPool() {
	b.pool = "pool"
}

func main() {
	
	// b := make(WoodenHouseBuilder)
	
	woodBulder := &WoodenHouseBuilder{}
	
	d := NewDirector()
	d.NewWoodenHouse(woodBulder)
	fmt.Println(*woodBulder) // {walls wooden roof pool}

	stoneBulder := &StoneHouseBuilder{}
	d.NewStoneHouse(stoneBulder)
	fmt.Println(*stoneBulder) // {wooden walls roof}
	
	
	
	// d := newDirector(b)
	// d.Make()
	
	// woodenHouse := b.GetResult()
}

// нужно построить два дома: каменный и деревянный с бассейном