package main

import "fmt"

func main() {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i + 42
	}

	for _, v := range arr {
		fmt.Printf("%v\t%T\n", v, v)
	}
	fmt.Printf("%T \t %v\n", arr, arr)
	fmt.Printf("%T \t %#v\n", arr, arr)
}

/*
Using a COMPOSITE LITERAL:
create a SLICE of TYPE int
assign these 10 VALUES
42, 43, 44, 45, 46, 47, 48, 49, 50, 51
Range over the slice and print the values out.
print out the VALUE and the TYPE
*/
