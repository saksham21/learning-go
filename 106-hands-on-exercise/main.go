package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4}
	for _, v := range arr {
		fmt.Printf("%v\t%T\n", v, v)
	}
	fmt.Println("-------")
	arr1 := [5]int{}
	for i := 0; i < 5; i++ {
		arr1[i] = i
	}

	for _, v := range arr1 {
		fmt.Printf("%v\t%T\n", v, v)
	}
}

/*
Using a COMPOSITE LITERAL:
create an ARRAY which holds 5 VALUES of TYPE int
assign VALUES to each index position.
Range over the array and print the values out.
print out the VALUE and the TYPE
*/
