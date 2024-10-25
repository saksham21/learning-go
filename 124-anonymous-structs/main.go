package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// type engine struct {
// 	engineModel string
// 	hp          int
// }

// func (e engine) start() {
// 	fmt.Println("Engine started. ", e.engineModel, e.hp)
// }

// type Car struct {
// 	engine
// 	carModel string
// 	brand    string
// }

// func (c Car) start() {
// 	fmt.Println("Car started. ", c.Engine, c.carModel, c.brand)
// }

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Tommy",
		last:  "Tom",
		age:   42,
	}
	fmt.Printf("%T\t%v\t%#v\n", p1, p1, p1)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Printf("%T\t%v\t%#v\n", p2, p2, p2)
	fmt.Println(p2.first, p2.last, p2.age)

	// engine := engine{
	// 	engineModel: "TurboBoost",
	// 	hp:          1000,
	// }
	// car := Car{
	// 	engine:   engine,
	// 	carModel: "A5",
	// 	brand:    "Audi",
	// }
	// car.start()
}
