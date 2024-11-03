package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	First string
	Last  string
	Age   int
}

func main() {
	u1 := user{First: "Todd", Last: "McLead", Age: 32}
	u2 := user{First: "James", Last: "Bond", Age: 33}
	u3 := user{First: "Jenny", Last: "MoneyPenny", Age: 23}

	ui := []user{u1, u2, u3}
	by, err := json.Marshal(ui)
	if err != nil {
		log.Fatalf("Err in json marshal: %s", err)
	}
	fmt.Println(by)
	fmt.Println(string(by))
}
