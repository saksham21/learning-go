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
		last:       "MoneyPenny",
		favFlavors: []string{"vanilla", "strawberry", "belgium chocolate"},
	}

	// fmt.Printf("%T\n", p1)
	// fmt.Printf("%#v\n", p1)
	// for _, v := range p1.favFlavors {
	// 	fmt.Printf("%v, ", v)
	// }
	// fmt.Println()
	// fmt.Println("-----------------")
	// fmt.Printf("%T\n", p2)
	// fmt.Printf("%#v\n", p2)
	// for _, v := range p2.favFlavors {
	// 	fmt.Printf("%v ", v)
	// }
	// fmt.Println()
	// fmt.Println("-----------------")

	mp := make(map[string]person)
	mp[p1.last] = p1
	mp[p2.last] = p2

	for _, v := range mp {
		fmt.Printf("%T\t%v\n", v, v)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-----------------")
	}
}

/*
Take the code from the previous exercise, then store the VALUES of type person in a map with the KEY of last name. Access each value in the map. Print out the values, ranging over the slice.
*/
