package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	xs := []string{"random", "rainbow", "delights", "in", "torpedo"}
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
