package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// unfurling a slice
	fmt.Println(sum(xi...))
}

// variadic parameter should be the final incoming parameter
func sum(a ...int) int {
	fmt.Printf("%T\t%v\n", a, a)
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}
