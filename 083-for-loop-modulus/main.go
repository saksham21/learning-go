package main

import (
	"fmt"
)

func main() {
	var i int
	for i < 20 {
		if i&1 == 1 {
			fmt.Printf("Value of i is %v\n", i)
		}
		i++
	}

	fmt.Println()
}
