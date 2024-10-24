package main

import (
	"fmt"
)

func main() {
	a := []string{"Harry", "Potter", "Magic", "University"}
	b := []string{"David", "Miller", "Normal", "University"}
	fmt.Println(a)
	fmt.Println(b)

	c := [][]string{a, b}
	fmt.Println(c)

	d := make([][]int, 3)
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d[0]))

}
