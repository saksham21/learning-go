package main

import "fmt"

func main() {
	mp := make(map[string][]string)
	mp["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	mp["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	mp["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	for _, v := range mp {
		for i, val := range v {
			fmt.Printf("%v\t val: %v\n", i, val)
		}
		fmt.Println("-----------------")
	}
}
