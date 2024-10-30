package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, "and my age is", p.age)
}

func main() {
	p1 := person{
		first: "Todd",
		age:   42,
	}
	p1.speak()
}

/*
Create a user defined struct with
the identifier “person”
the fields:
first
age
attach a method to type person with
the identifier “speak”
the method should have the person say their name and age
create a value of type person
call the method from the value of type person
*/
