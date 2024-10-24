package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("a - %#v\n", a)
	fmt.Println("------------")

	// [inclusive: exclusive]
	fmt.Printf("a - %#v\n", a[0:4])
	fmt.Println("------------")

	fmt.Printf("a - %#v\n", a[:7])
	fmt.Println("------------")

	fmt.Printf("a - %#v\n", a[4:])
	fmt.Println("------------")

	fmt.Printf("a - %#v\n", a[:])
	fmt.Println("------------")

	fmt.Printf("a - %#v\n", a[len(a):])
	fmt.Println("------------")
}
