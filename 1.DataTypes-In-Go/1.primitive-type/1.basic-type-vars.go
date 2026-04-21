/*
	Basic Type Variables in Go
	===========================
	In Go, you can declare variables of basic data types such as string, int, float64, and bool.
	You can use the `var` keyword to declare a variable and specify its type. The variable can then
	be assigned a value of the specified type.

	In the example below, we declare variables of different types and assign them values. We then
	print the values of these variables using the `fmt.Println` function.

	Different Primitive Types:
	- string: A sequence of characters
	- int: 
		- int8, int16, int32, int64 [signed integers]
		- uint8, uint16, uint32, uint64 [unsigned integers]
	- float64: 
		- float32, float64 [floating-point numbers]
	- bool: A boolean value
	- rune: A Unicode code point (alias for int32)
	- byte: An alias for uint8
*/
package main

import "fmt"

func main() {
	// Declare variables of different types
	var name string = "Go Programming"
	var age int = 10
	var salary float64 = 50000.50
	var isFun bool = true

	// Print the values of the variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Is Go fun?", isFun)


	// multiple variable declaration
	var (
		name2 string = "Go Programming"
		age2  int    = 10
	)
	fmt.Println("Name:", name2)
	fmt.Println("Age:", age2)

	// short variable declaration
	message, customMessage := "Hello", "World"
	fmt.Println(message, customMessage)
}
