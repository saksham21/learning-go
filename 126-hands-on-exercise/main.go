package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first:      "Todd",
		last:       "Depp",
		favFlavors: []string{"chocolate", "strawberry", "casatta"},
	}

	p2 := person{
		first:      "Jonny",
		last:       "Depp",
		favFlavors: []string{"vanilla", "strawberry", "belgium chocolate"},
	}

	fmt.Printf("%T\n", p1)
	fmt.Printf("%#v\n", p1)
	for _, v := range p1.favFlavors {
		fmt.Printf("%v, ", v)
	}
	fmt.Println()
	fmt.Println("-----------------")
	fmt.Printf("%T\n", p2)
	fmt.Printf("%#v\n", p2)
	for _, v := range p2.favFlavors {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

*/
