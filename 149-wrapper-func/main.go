package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("Error in main: %s\n", err)
	}
	fmt.Println(f)
	fmt.Println(string(f))
}

func readFile(fn string) ([]byte, error) {
	f, err := os.ReadFile(fn)
	if err != nil {
		return nil, fmt.Errorf("error in reading file: %s", err)
	}
	return f, nil
}
