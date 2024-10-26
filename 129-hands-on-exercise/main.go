package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Friend1": 1,
			"Friend2": 2,
		},
		favDrinks: []string{"coke", "beer", "lassi"},
	}

	fmt.Printf("%T\t%v\n", p1, p1)
	fmt.Println(p1.first, p1.friends, p1.favDrinks)

	for k, v := range p1.friends {
		fmt.Println(p1.first, "- friends -", k, v)
	}
	for _, v := range p1.favDrinks {
		fmt.Println(p1.first, "- drinks -", v)
	}
}

/*
Create and use an anonymous struct with these fields:
first     string
friends   map[string]int
favDrinks []string
Print things.
*/
