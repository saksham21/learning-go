package main

import "fmt"

// value semantics
func addOne(n int) int {
	n += 1
	return n
}

// pointer semantics
func addOnePointer(n *int) {
	*n += 1
}

func main() {
	// value semantics
	a := 1
	fmt.Println(a)         // 1
	fmt.Println(addOne(a)) // 2
	fmt.Println(a)         // 1

	// pointer semantics
	addOnePointer(&a)
	fmt.Println(a) // 2
}
