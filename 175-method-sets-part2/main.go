package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "I'm walking")
}
func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "I'm running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{
		first: "Henry",
	}
	d1.walk()
	d1.run()

	// this wont work since d1 does not implement run method, so it does not belong to interface
	// func (d *dog) run() --> func (d dog) run() ::>> this will resolve for both value and pointer semantics
	// youngRun(d1)

	d2 := &dog{
		first: "Padget",
	}
	d2.walk()
	d2.run()

	// this will work since d2 does implement run method and walk method, so it does belong to interface
	youngRun(d2)
}
