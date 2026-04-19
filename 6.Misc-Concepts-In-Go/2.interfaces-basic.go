package main

import (
	"fmt"
	"math"
)

// define an inteface
type Shape interface{
	Area() float64
}


// struct-1: Rectangle
type Rectangle struct {
	width float64
	height float64
}

// struct-2 Circle
type Circle struct {
	radius float64
}


// functions to write
func (r Rectangle) Area() float64 {
	return  r.width * r.height
}

func (c Circle) Area() float64 {
	return  math.Pow(c.radius, 2)
}

func printArea(s Shape){
	fmt.Println("Area: ", s.Area())
}

func main(){

	// instantiated
	rectangle := Rectangle{ width: 10.12, height: 20.23}
	circle := Circle{ radius: 3.23}

	printArea(rectangle)
	printArea(circle)
}
