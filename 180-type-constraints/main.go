package main

import "fmt"

func main() {
	fmt.Println(addI(2, 3))
	fmt.Println(addF(2.2, 3.2))

	fmt.Println(addT(2, 3))
	fmt.Println(addT(2.2, 3.2))
}

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

func addT[T int | float64 | string](a, b T) T {
	return a + b
}
