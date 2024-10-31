package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)
	intDelta1(a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta1(xi)
	fmt.Println(xi)

	sliceDelta(&xi)
	fmt.Println(xi)

	mp := make(map[int]int)
	mp[1] = 1
	mp[2] = 2
	fmt.Println(mp)
	sliceMap(mp, 3)
	fmt.Println(mp)
}

func sliceMap(n map[int]int, i int) {
	n[i] = i
}

func sliceDelta(ni *[]int) {
	(*ni)[0] = 101
	*ni = append(*ni, 0)
}

func sliceDelta1(ni []int) {
	ni[0] = 100
	ni = append(ni, 99)
}

func intDelta(n *int) {
	fmt.Println(" ----- ", n)
	*n = 43
}

func intDelta1(n int) {
	fmt.Println(" ----- ", &n)
	a := &n
	*a = 44
}
