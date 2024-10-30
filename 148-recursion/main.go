package main

import (
	"fmt"
	"time"
)

func main() {

	// fmt.Println(fact(4))
	fmt.Println(factLoop(4))
	getTime(fact)
}

func getTime(f func(int) int) {
	t := time.Now()
	fmt.Println(f(4))
	fmt.Println(time.Since(t))
}

func fact(i int) int {
	// time.Sleep(2 * time.Second)
	if i == 0 {
		return 1
	}
	return i * fact(i-1)
}

func factLoop(i int) int {
	ans := 1
	for i >= 1 {
		ans *= i
		i--
	}
	return ans
}
