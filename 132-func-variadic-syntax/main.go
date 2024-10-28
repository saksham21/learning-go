package main

import "fmt"

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(sum())
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
