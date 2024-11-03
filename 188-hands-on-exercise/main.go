package main

import (
	"fmt"
	"sort"
)

type user struct {
	First string
	Last  string
	Age   int
}

type byAge []user

func (ul byAge) Len() int           { return len(ul) }
func (ul byAge) Less(i, j int) bool { return ul[i].Age < ul[j].Age }
func (ul byAge) Swap(i, j int)      { ul[i], ul[j] = ul[j], ul[i] }

type byLast []user

func (ul byLast) Len() int           { return len(ul) }
func (ul byLast) Less(i, j int) bool { return ul[i].Last < ul[j].Last }
func (ul byLast) Swap(i, j int)      { ul[i], ul[j] = ul[j], ul[i] }

func main() {
	u1 := user{First: "Todd", Last: "McLead", Age: 32}
	u2 := user{First: "James", Last: "Bond", Age: 33}
	u3 := user{First: "Jenny", Last: "MoneyPenny", Age: 23}

	users := []user{u1, u2, u3}
	fmt.Println(users)
	fmt.Println("After sorting by age: ")
	sort.Sort(byAge(users))
	fmt.Println(users)

	fmt.Println("---------------")
	fmt.Println("After sorting by last: ")
	sort.Sort(byLast(users))
	fmt.Println(users)
}
