package main

import (
	"fmt"
)

func main() {
	fmt.Println(addI(2, 3))
	fmt.Println(addF(2.2, 3.2))

	var a tInt = 2
	var b tInt = 2
	var c tFloat64 = 2.2
	var d tFloat64 = 3.2
	var e tString = "4"
	var f tString = "2"
	fmt.Println(addT(a, b))
	fmt.Println(addT(c, d))
	fmt.Println(addT(e, f))
}

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type tInt int

func (t tInt) print() {
	fmt.Println("Printing int:", t)
}

type tFloat64 float64

func (t tFloat64) print() {
	fmt.Println("Printing float64:", t)
}

type tString string

func (t tString) print() {
	fmt.Println("Printing string:", t)
}

// underlying types by using ~
type myTypes interface {
	~int | ~float64 | ~string
	print()
}

// option 1 create new interface and add print method there

type myTypes2 interface {
	print()
}

func print(m myTypes2) {
	m.print()
}

// option 2 is just call print method directly
// a.print()
// b.print()

func addT[T myTypes](a, b T) T {
	print(a)
	print(b)
	return a + b
}
