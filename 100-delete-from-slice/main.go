package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("a - %#v\n", a)
	fmt.Println("------------")

	// delete element 4
	a = append(a[:4], a[5:]...)
	fmt.Printf("a - %#v\n", a)
	fmt.Println("------------")
}
