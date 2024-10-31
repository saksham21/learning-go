package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%T\t%T\n", x, &x)
	y := "Hello"
	fmt.Println(y)
	fmt.Println(&y)
	fmt.Printf("%T\t%T\n", y, &y)
}
