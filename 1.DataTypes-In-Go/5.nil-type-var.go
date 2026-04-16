/*
	Nil Type Var in Go
	========================
	In Go, the `nil` value represents the zero value for pointers, interfaces, maps, slices, channels, and function types. It indicates that a variable does not point to any valid data or object. When you declare a variable of one of these types without initializing it, it will have the `nil` value by default.

	In the example below, we declare a variable named `nilValue` of type `interface{}` and assign it the `nil` value. This means that `nilValue` does not point to any valid data or object.
*/

package main

import "fmt"

func main() {
	var nilValue interface{} = nil // The 'nil' value can be assigned to an interface type
	fmt.Println("The value of nilValue is:", nilValue)
}
