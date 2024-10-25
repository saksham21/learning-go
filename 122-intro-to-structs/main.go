package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Printf("%T\t%v\t%#v\n", p1, p1, p1)
	fmt.Println(p1.first, p1.last, p1.age)
}
