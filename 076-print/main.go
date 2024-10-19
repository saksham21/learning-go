package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4}
	b := make(map[string]int)
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i)
	}
	b["one"] = 1
	b["two"] = 2
	fmt.Println()
	fmt.Println(a)
	fmt.Println(b)
}
