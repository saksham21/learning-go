package main

import "fmt"

func main() {
	// arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i + 42
	}
	fmt.Println(arr)
	arr = append(arr, 52)
	fmt.Println(arr)
	arr = append(arr, 53, 54, 55)
	fmt.Println(arr)
	y := []int{56, 57, 58, 59, 60}
	arr = append(arr, y...)
	fmt.Println(arr)
}

/*

Follow these steps:
start with this slice
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
append to that slice this value
52
print out the slice
in ONE STATEMENT append to that slice these values
53
54
55
print out the slice
append to the slice this slice
y := []int{56, 57, 58, 59, 60}
print out the slice

*/