package main

import "fmt"

func main() {

	a := doMath(42, 22, add)
	b := doMath(42, 22, subtract)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

}

func doMath(i, j int, f func(int, int) int) int {
	return f(i, j)
}

func add(i, j int) int {
	return i + j
}

func subtract(i, j int) int {
	return i - j
}
