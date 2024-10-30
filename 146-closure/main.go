package main

import "fmt"

func main() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("-----------")
	f1 := incrementor()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
