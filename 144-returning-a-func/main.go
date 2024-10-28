package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	y := bar()
	fmt.Println(y())

	z := foobar()
	fmt.Println(z(2, 2))

	fmt.Println("-----------")
	fmt.Printf("%T\t%T\t%T\n", x, y, z)
	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", foobar)

}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}

// anon func
// func(p parameters(s)) (r return(s)) {code}
func foobar() func(i, j int) int {
	return func(i, j int) int {
		return 45 + i*j
	}
}
