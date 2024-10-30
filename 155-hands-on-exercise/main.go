package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
	width  float64
}

func (s Square) Area() float64 {
	return s.length * s.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Shape interface {
	Area() float64
}

func Info(s Shape) {
	fmt.Println("Area is", s.Area())
}

func main() {
	s := Square{
		length: 10,
		width:  20,
	}
	Info(s)
	fmt.Printf("%T\n", s)

	c := Circle{
		radius: 10,
	}
	Info(c)
	fmt.Printf("%T\n", c)
}

/*
create a type SQUARE
length int
width int
create a type CIRCLE
radius float64
attach a method to each that calculates AREA and returns it
circle area= π r 2
math.Pi
math.Pow
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle*/
