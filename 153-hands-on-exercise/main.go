package main

import "fmt"

func main() {
	defer fmt.Println("main")
	defer func() {
		fmt.Println("Func1")
	}()
	defer func() {
		fmt.Println("Func2")

	}()
	defer func() {
		fmt.Println("Func3")

	}()
}

/*

“defer” multiple functions in main
show that a deferred func runs after the func containing it exits.
determine the order in which the multiple defer funcs run

*/

// deferred functions run in LIFO order
// last in first out LIFO
