package main

import "fmt"

// func (r receiver) (p parameter(s)) (r return(s)) { code }
func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
