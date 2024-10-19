package main

import (
	"fmt"
)

func main() {
	var i int = 5
	for i > 0 {
		fmt.Printf("Value of i is %v\n", i)
		i--
	}

	fmt.Println()

	for i < 5 {
		fmt.Printf("Value of i is %v\n", i)
		i++
	}
}
