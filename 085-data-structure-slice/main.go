package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47}

	for i, j := range x {
		fmt.Printf("Value for Index %v is %v\n", i, j)
	}
}
