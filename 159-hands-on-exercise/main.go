package main

import "fmt"

func main() {
	fmt.Println(Paradise("India"))
}

/*
Paradise prints our the user's input of paradise as a sentence
Hello World
Hello Gophers
*/
// Heleloeloeleo
func Paradise(s string) string {
	return fmt.Sprint("My paradise is ", s)
}

// to run your test
// at the terminal, aka bash, aka shell
// go test
