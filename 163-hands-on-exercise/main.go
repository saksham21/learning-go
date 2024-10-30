package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func foo() func() int {
	x := 42
	return func() int {
		x -= 10
		return x
	}
}

// func return
