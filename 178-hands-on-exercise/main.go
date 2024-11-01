package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	// d.first = "Sammy"
	fmt.Println("Dog", d.first, "is walking")
}
func (d dog) run() {
	fmt.Println("Dog", d.first, "is running")
}

type youngin interface {
	walk()
	run()
}

func youngWalk(y youngin) {
	y.walk()
}

func youngRun(y youngin) {
	y.run()
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}

func (d *dog) changeNameByPointer(s string) {
	d.first = s
}

func main() {
	d1 := dog{
		first: "Todd",
	}
	youngRun(d1)
	youngWalk(d1)

	fmt.Println(d1)
	d2 := &dog{
		first: "Teddy",
	}
	youngRun(d2)
	youngWalk(d2)
	fmt.Println(d2)

	fmt.Println("Changing Dog Names")
	d1 = d1.changeName("Sammy007")
	fmt.Println(d1)

	d2.changeNameByPointer("Sammy123")
	fmt.Println(*d2)

}
