package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(s string, i int) {
		fmt.Println("Hi", s, ", This is an anon func from year", i)
	}("Saksham", 2024)
}

func foo() {
	fmt.Println("Foo ran")
}

// a named function with an indentifier
// func (r reciever) indentifier(p parameters(s)) (r return(s)) { code }

// an anonymous function
// func (p parameters(s)) (r return(s)) { code }
