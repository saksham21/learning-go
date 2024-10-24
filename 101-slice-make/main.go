package main

import (
	"fmt"
)

func main() {
	a := []string{"a", "b", "c"}
	fmt.Println(a)

	// creating slice with make
	b := make([]int, 0, 10)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	// creating slice with new
	c := new([10]int)[:5]
	fmt.Printf("%#v\n", c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
}
