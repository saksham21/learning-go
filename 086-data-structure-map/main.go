package main

import "fmt"

func main() {
	x := map[string]int{"James": 42, "Moneypenny": 32}

	for key, value := range x {
		fmt.Printf("Value for key %v is %v\n", key, value)
	}

	fmt.Println(main1(11))
}

func main1(i int) int {
	return i + 1
}
