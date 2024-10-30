package main

import "fmt"

func main() {
	f1 := func() {
		fmt.Println("Hello F1")
	}
	f1()
	f2 := func(a int) {
		fmt.Println("Hello F1", a)
	}
	f2(42)
}

// func expression
