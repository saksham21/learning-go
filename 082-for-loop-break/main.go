package main

import (
	"fmt"
)

func main() {
	var i int
	for {
		if i >= 5 {
			break
		}
		fmt.Printf("Value of i is %v\n", i)
		i++
	}

	fmt.Println()

	for {
		if i < 0 {
			break
		}
		fmt.Printf("Value of i is %v\n", i)
		i--
	}
}
