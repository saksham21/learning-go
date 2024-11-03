package main

import (
	"fmt"
)

func main() {
	fmt.Println(addI(2, 3))
	fmt.Println(addF(2.2, 3.2))

	var a int = 2
	var b int = 2
	fmt.Println(addT(a, b))

	var c float64 = 2.2
	var d float64 = 3.2
	fmt.Println(addT(c, d))

	var e string = "4"
	var f string = "2"
	fmt.Println(addT(e, f))
}

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myTypes interface {
	int | float64 | string
}

func addT[T myTypes](a, b T) T {
	return a + b
}
