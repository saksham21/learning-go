package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%T\t%T\n", x, &x)
	y := &x
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Println(*y)
	fmt.Println(*&x)

	*y = 43
	fmt.Println(x, y, &x, *y)

}
