package main

import "fmt"

func main() {
	foo()
	bar("Todd")

	a := foobar("Jenny")
	fmt.Println(a)

	b, c := aloha("James", "Beer")
	fmt.Println(b, c)
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println(s, " enjoys martini in bar")
}

func foobar(s string) string {
	return fmt.Sprint(s, " enjoys martini in foobar")
}

func aloha(name string, drink string) (string, string) {
	return fmt.Sprint(name, " enjoys ", drink, " in foobar"), "aloha"
}

// func (r reciever) identifier (p parameters(s)) (return(s)) { code }
/*
func syntax

 no params, no returns
 1 params, no returns
 1 params, 1 returns
 2 params, 2 returns
*/
