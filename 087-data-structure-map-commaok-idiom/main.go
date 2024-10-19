package main

import "fmt"

func main() {
	x := map[string]int{"James": 42, "Moneypenny": 32}

	for key, value := range x {
		fmt.Printf("Value for key %v is %v\n", key, value)
	}

	fmt.Println(x["James"])
	fmt.Println(x["Jame"])

	// comma ok idiom
	if v, ok := x["James"]; ok {
		fmt.Println("Value of James is ", v)
	} else {
		fmt.Println("Value of James NOT FOUND")
	}
	if v, ok := x["Jame"]; ok {
		fmt.Println("Value of Jame is ", v)
	} else {
		fmt.Println("Value of Jame NOT FOUND")
	}
	keys := make([]string, 0, len(x))
	for k := range x {
		keys = append(keys, k)
	}
	fmt.Println("Printing keys ----->", keys)

	values := make([]int, 0, len((x)))
	for _, v := range x {
		values = append(values, v)
	}
	fmt.Println("Printing values ----->", values)
}
