package main

import "fmt"

func main() {
	// arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i + 42
	}
	fmt.Println(arr)
	arr = append(arr[:3], arr[6:]...)
	fmt.Println(arr)
}

/*
To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
start with this slice
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
use APPEND & SLICING to get these values here which you should ASSIGN to the variable “x” and then print:
[42, 43, 44, 48, 49, 50, 51]

*/
