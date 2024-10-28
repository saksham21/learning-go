package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("My name is", p.first)
}

// func (r receiver) (p parameter(s)) (r return(s)) { code }
func main() {

	p1 := person{first: "James"}
	p2 := person{first: "Jenny"}

	p1.speak()
	p2.speak()
}
