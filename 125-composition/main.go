package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type foo int

func (a foo) dance(b foo) bool {
	return a > b
}

func main() {
	var a foo = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b := a.dance(41)
	fmt.Println(b)

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Tommy",
		last:  "Tom",
		age:   42,
	}
	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
}
