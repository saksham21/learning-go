package main

import "fmt"

func main() {
	fmt.Println("Total is: ", sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func sum(ii []int) (total int) {
	total = 0
	for _, i := range ii {
		total += i
	}
	return total
}

/*
named return vs normal return (unnamed return)
*/
