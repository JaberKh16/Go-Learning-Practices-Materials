/*
	Interfaces Concepts In Go
	=========================

	# What Are Interfaces?

	Interfaces in Go define a set of method signatures (behavior).
	A type satisfies an interface automatically if it implements those methods.

	There is NO explicit "implements" keyword in Go.

	# Basic Interface Syntax

	You define an interface using the "type" keyword:

		type Shape interface {
			Area() float64
		}

	Any type that has an Area() method automatically satisfies Shape.

	# Implementing an Interface

	Example:

		type Rectangle struct {
			width, height float64
		}

		func (r Rectangle) Area() float64 {
			return r.width * r.height
		}

	Now Rectangle satisfies the Shape interface.

	# Multiple Types, One Interface

	Different types can implement the same interface:

		type Circle struct {
			radius float64
		}

		func (c Circle) Area() float64 {
			return 3.14 * c.radius * c.radius
		}

	Both Rectangle and Circle are considered Shape.

	# Using Interfaces as Parameters

	Interfaces allow flexible function design:

		func printArea(s Shape) {
			fmt.Println(s.Area())
		}

	You can pass any type that satisfies Shape.

	# Empty Interface (any)

	The empty interface allows any type:

		func printValue(v any) {
			fmt.Println(v)
		}

	"any" is an alias for interface{}.

	# Type Assertion

	You can extract the underlying value from an interface:

		value := v.(int)

	This will panic if the type is incorrect.
	Safe version:

		value, ok := v.(int)

	# Type Switch

	You can check multiple types:

		switch val := v.(type) {
		case int:
			fmt.Println("int:", val)
		case string:
			fmt.Println("string:", val)
		default:
			fmt.Println("unknown type")
		}

	# Why Use Interfaces?

	- Enables abstraction
	- Allows flexible and reusable code
	- Helps decouple components
	- Useful for testing (mock implementations)

	# Key Idea

	Interfaces define behavior, not data.
	If a type implements the methods, it satisfies the interface automatically.
*/

// program demonstrates embedding interfaces
package main

import (
	"fmt"
	"math"
)

// define an inteface
type Shape interface{
	Area() float64
}

type Measurable interface {
	Perimeter() float64
}


// embeded interface
type Geometry interface {
	Shape
	Measurable
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

// struct-3 Error
type CalculationError struct {
	message string
}


// functions to write
func (r Rectangle) Perimeter() float64 {
	return  2 * (r.width + r.height) 
}


func (r Rectangle) Area() float64 {
	return  r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Perimeter() float64 {
	return  2 * math.Pi * c.radius
}


func descibeStruct(t interface {}){
	fmt.Println("Structs: ", t)
}

func descibeShape(g Geometry) float64 {
	fmt.Println("Area: ", g.Area())
	fmt.Println("Perimeter: ", g.Perimeter())
}

func (ce CalculationError) Error() string {
	return ce.message
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return  0, CalculationError{message: "value can't be zero"}
	}
	return  math.Sqrt(val), nil
}


func main(){

	// instantiated
	rectangle := Rectangle{ width: 10.12, height: 20.23}
	circle := Circle{ radius: 3.23}

	// see the structs
	descibeStruct(rectangle)
	descibeStruct(circle)


	// calculate the area
	descibeShape(rectangle)
	descibeShape(circle)
}
