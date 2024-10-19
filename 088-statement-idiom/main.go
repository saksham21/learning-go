package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var c int = 1
	for i := 0; i < 100; i++ {
		// statement statement idiom
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("Total count is %v\t Value of x is %v at iteration %v\n", c, x, i)
			c++
		}
	}
}
