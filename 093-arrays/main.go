package main

import "fmt"

func main() {
	var x [5]int
	fmt.Printf("%v\t\t\t%T\n", x, x)
	fmt.Println(len(x))
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	y := [...]int{0, 0, 0, 0, 0}
	fmt.Println(y)
	y = x
	fmt.Println(y)
}
