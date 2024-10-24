package main

import "fmt"

func main() {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i + 42
	}

	// for _, v := range arr {
	// 	fmt.Printf("%v\t%T\n", v, v)
	// }
	// fmt.Printf("%T \t %v\n", arr, arr)
	// fmt.Printf("%T \t %#v\n", arr, arr)

	fmt.Println(arr[:5])
	fmt.Println(arr[5:])
	fmt.Println(arr[2:7])
	fmt.Println(arr[1:6])
}

/*
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]

*/
