package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	i := "This is my world."
	j := "Hello World."
	k := "My name is Saksham."
	x := 42

	a = &i
	b = &j
	c = &k
	d = &x
}

func main() {

	fmt.Printf("Type: %T\t Address: %v\t Value: %v\n", a, a, *a)
	fmt.Printf("Type: %T\t Address: %v\t Value: %v\n", b, b, *b)
	fmt.Printf("Type: %T\t Address: %v\t Value: %v\n", c, c, *c)
	fmt.Printf("Type: %T\t Address: %v\t Value: %v\n", d, d, *d)
}
