package main

import (
	"fmt"
	"time"
)

func main() {
	timeFun(fun1)
	timeFun(fun)
}

// this is a wrapper function where we passed f func() as a callback as an argument
func timeFun(f func()) {
	t := time.Now()
	f()
	fmt.Println("Time elapsed:", time.Since(t))
}

func fun() {
	fmt.Println("Goroutine going to sleep")
	time.Sleep(1 * time.Second)
	fmt.Println("Goroutine wake up from sleep")
}

func fun1() {
	for i := 0; i < 2000; i++ {
		fmt.Println(i)
	}
}

// wrapper
