package main

import "fmt"

func main() {
	xxs := make([][]string, 0)
	xxs = append(xxs, []string{"James", "Bond", "Shaken, not stirred"})
	xxs = append(xxs, []string{"Miss", "Moneypenny", "I'm 008."})

	for i, xs := range xxs {
		fmt.Printf("For i=%v\n", i)
		for j, v := range xs {
			fmt.Printf("j=%v val=%v\n", j, v)
		}
		fmt.Println("-------")
	}
}

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."

Range over the records, then range over the data in each record.

*/
