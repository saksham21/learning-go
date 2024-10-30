package main

import "fmt"

func main() {
	fmt.Println(paradise("India"))
}

func paradise(s string) string {
	return fmt.Sprint("My paradise is ", s)
}

// to run your test
// at the terminal, aka bash, aka shell
// go test
