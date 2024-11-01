package main

import "fmt"

type person struct {
	first string
}

func someFunc1(p person, s string) person {
	p.first = s
	return p
}

func someFunc2(p *person, s string) {
	p.first = s
	fmt.Println("someFunc2 called")
}

func main() {
	p1 := person{first: "Todd"}
	p1 = someFunc1(p1, "Sammy007")
	fmt.Println(p1)

	someFunc2(&p1, "Sammy124")
	fmt.Println(p1)
}
