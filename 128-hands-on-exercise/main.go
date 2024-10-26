package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	e1 := engine{
		electric: false,
	}
	e2 := engine{
		electric: true,
	}
	v1 := vehicle{
		engine: e1,
		make:   "audi",
		model:  "A5",
		doors:  4,
		color:  "Red",
	}
	v2 := vehicle{
		engine: e2,
		make:   "bmw",
		model:  "E3",
		doors:  2,
		color:  "Black",
	}

	fmt.Printf("%T\t%v\n", v1, v1)
	fmt.Printf("%T\t%v\n", v2, v2)
	fmt.Println(v1.engine, v1.electric, v1.make, v1.model, v1.doors, v1.color)
	fmt.Println(v2.engine, v2.electric, v2.make, v2.model, v2.doors, v2.color)
	// fmt.Println(v1.engine, v1.engine.electric, v1.make, v1.model, v1.doors, v1.color)
	// fmt.Println(v2.engine, v2.engine.electric, v2.make, v2.model, v2.doors, v2.color)
}

/*
Create a type engine struct, and include this field
electric bool
Create a type vehicle struct, and include these fields
engine
make
model
doors
color
Create two VALUES of TYPE vehicle
use a composite literal
Print out each of these values.
Print out a single field from each of these values.

*/
