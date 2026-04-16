/*
	Const Type Var in Go
	========================
	In Go, you can declare constants using the `const` keyword. Constants are immutable values that cannot
	be changed after they are defined. They can be of any basic data type, such as string, int, float64,
	or bool.

	In the example below, we declare a constant named `name` of type `string` and assign it the value
	"Go Programming". Since `name` is a constant, attempting to reassign it will result in
	a compile-time error.
*/

package main

import "fmt"

func main() {
	const name string = "Go Programming"
	// name := "New Name" // This will cause a compile-time error because 'name' is a constant

	fmt.Println("Name:", name)
}