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
}
