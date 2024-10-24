package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println("------------")
	a = append(a, 5)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println("------------")
	b := make([]int, 0, 10)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	b = append(b, a...)
	b = append(b, a...)
	b = append(b, a...)
	b = append(b, a...)
	fmt.Println("------------")
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	// cap is increased by 2 * prev cap every time len is increased

	b = append(b[0:5], 100)
	fmt.Println("------------")
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
}
