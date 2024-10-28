package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

type count int

func (b book) String() string {
	return fmt.Sprint("The title of book is ", b.title)
}

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func printLog(a fmt.Stringer) {
	log.Println("::: LOG FROM 138 :::", a.String())
}

func main() {
	b1 := book{
		title: "West With The Night",
	}

	var n count = 42
	printLog(b1)
	printLog(n)
}
