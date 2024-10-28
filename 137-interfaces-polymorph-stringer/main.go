package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

type count int
type count1 int

func (b book) String() string {
	return fmt.Sprint("The title of book is ", b.title)
}

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func main() {
	b1 := book{
		title: "West With The Night",
	}

	var n count = 42
	var m count1 = 42
	fmt.Println(b1)
	fmt.Println(n)
	fmt.Println(m)
}
