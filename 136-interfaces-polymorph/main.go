package main

import "fmt"

type person struct {
	first string
}
type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("My name is", p.first)
}
func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

// person and secretagent both got speak() method, so they are of type=human
type human interface {
	speak()
	// write()
}

func saySomething(h human) {
	h.speak()
}

// func (r receiver) (p parameter(s)) (r return(s)) { code }
func main() {

	sa1 := secretAgent{
		person: person{first: "James"},
		ltk:    true,
	}
	p2 := person{first: "Jenny"}

	// sa1.speak()
	// p2.speak()

	saySomething(sa1)
	saySomething(p2)
}
