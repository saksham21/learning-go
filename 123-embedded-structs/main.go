package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretagent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Printf("%T\t%v\t%#v\n", p1, p1, p1)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println("--------------")
	sa := secretagent{
		person: p1,
		ltk:    true,
	}

	fmt.Printf("%T\t%v\t%#v\n", sa, sa, sa)
	fmt.Println(sa.first, sa.last, sa.age, sa.ltk, sa.person)
}
