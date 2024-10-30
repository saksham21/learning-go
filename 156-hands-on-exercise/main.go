package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Add(3, 4))
}

// to run your test
// at the terminal, aka bash, aka shell
// go test
