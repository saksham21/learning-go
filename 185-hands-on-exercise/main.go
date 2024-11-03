package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

// type userList []user

func main() {
	str := `[{"First":"Todd","Last":"McLead","Age":32},{"First":"James","Last":"Bond","Age":33},{"First":"Jenny","Last":"MoneyPenny","Age":23}]`

	var us []user

	err := json.Unmarshal([]byte(str), &us)
	if err != nil {
		log.Fatalf("Err in json unmarshal: %s", err)
	}
	fmt.Println(us)

	for i, person := range us {
		fmt.Println("Person #", i)
		fmt.Println(person.First, person.Last, person.Age)
	}
}
