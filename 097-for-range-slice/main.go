package main

import "fmt"

func main() {
	a := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}

	fmt.Printf("String is of Type %T and length %v\n", a, len(a))
	fmt.Println(a)

	for _, x := range a {
		fmt.Printf("%v\n", x)
	}
}
