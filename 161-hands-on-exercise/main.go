package main

import "fmt"

func main() {
	func(a, b int) {
		fmt.Println("Sum is", (a + b))
	}(10, 1)
}

// anon func
