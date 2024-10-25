package main

import "fmt"

func main() {
	// am := make(map[string]int)
	// am["Todd"]=42

	am := map[string]int{
		"Todd":   42,
		"Henry":  43,
		"Padget": 16,
	}
	fmt.Println(len(am))
	fmt.Printf("%v\n", am)
	fmt.Printf("%#v\n", am)

	for k, v := range am {
		fmt.Printf("key %v\t val: %v\n", k, v)
	}
	fmt.Println("------------------------")
	for _, v := range am {
		fmt.Printf("val: %v\n", v)
	}
	fmt.Println("------------------------")
	for k := range am {
		fmt.Printf("key: %v\n", k)
	}

	keys := make([]string, 0, len(am))
	for k := range am {
		keys = append(keys, k)
	}
	fmt.Printf("keys: %#v\n", keys)
	vals := make([]int, 0, len(am))
	for _, v := range am {
		vals = append(vals, v)
	}
	/*
		for k, v := range am {
			keys = append(keys, k)
			vals = append(vals, v)
		}
	*/
	fmt.Printf("vals: %#v\n", vals)
}
