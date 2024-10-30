package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	fmt.Println(doMath(42, 16, add))
	fmt.Println(doMath(42, 16, subtract))
}

// to run your test
// at the terminal, aka bash, aka shell
// go test
