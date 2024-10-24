package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	// b := a

	b := make([]int, len(a))
	copy(b, a)

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("-----------")

	a[0] = 7

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("-----------")

	c := make([]int, 0, len(a))
	c = append(c, a...)

	a[1] = 10

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("-----------")
}
