package main

import "fmt"

func main() {
	x := []string{"hello", "world"}
	fmt.Println(x)
	fmt.Printf("Type %T and len %v\n", x, len(x))
}
