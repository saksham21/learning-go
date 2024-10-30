package main

import "fmt"

func main() {
	x1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("foo", foo(x1...))
	fmt.Println("bar", bar(x1))
}

func foo(p ...int) int {
	sum := 0
	for _, i := range p {
		sum += i
	}
	return sum
	// return bar(p)
}

func bar(ii []int) int {
	sum := 0
	for _, i := range ii {
		sum += i
	}
	return sum
}

/*
create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in

create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in

*/
