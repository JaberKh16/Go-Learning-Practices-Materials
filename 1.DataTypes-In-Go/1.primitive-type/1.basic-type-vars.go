/*
Basic Type Variables in Go
===========================
In Go, you can declare variables of basic data types such as string, int, float64, and bool.
You can use the `var` keyword to declare a variable and specify its type. The variable can then
be assigned a value of the specified type.

In the example below, we declare variables of different types and assign them values. We then
print the values of these variables using the `fmt.Println` function.
*/
package main

import "fmt"

func main() {
	var name string = "Go Programming"
	var age int = 10
	var salary float64 = 50000.50
	var isFun bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Is Go fun?", isFun)
}