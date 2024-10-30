package main

import "fmt"

func main() {

	v1 := doMath(20, 10, add)
	v2 := doMath(20, 10, subtract)

	fmt.Println(v1)
	fmt.Println(v2)

}

func add(a, b int) int {
	return a + b
}
func subtract(a, b int) int {
	return a - b
}
func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

// callback
