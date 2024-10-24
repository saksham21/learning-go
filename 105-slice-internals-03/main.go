package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []float64{3, 1, 4, 2}
	b := make([]float64, len(a))
	// b := make([]float64, 0, len(a))
	// b = append(b, a...)
	fmt.Println(b)
	copy(b, a)

	fmt.Println(medianOne(a))
	fmt.Println("after medianOne", a)
	fmt.Println(medianTwo(b))
	fmt.Println("after medianTwo", b)
	fmt.Println("-----------")
}

func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if (len(x)%2)&1 == 1 {
		return x[i]
	} else {
		return (x[i-1] + x[i]) / 2
	}
}

func medianTwo(x []float64) float64 {
	y := make([]float64, len(x))
	copy(y, x)

	return medianOne(y)
}
