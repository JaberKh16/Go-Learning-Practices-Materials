/*
	User-Defined Functions in Go
	===========================
	In Go, you can define your own functions to perform specific tasks. A function is a reusable block of code that takes input parameters, performs a specific operation, and may return a value.

	To define a function in Go, you use the `func` keyword followed by the function name, a list of parameters (if any), and the return type (if any). The function body is enclosed in curly braces `{}`.

	In the example below, we define a simple function called `greet` that takes a string parameter `name` and returns a greeting message. We then call this function from the `main` function and print the result.
*/

package main

import "fmt"

// Function definition
func customMessage(name string) string {
	return "Hello, " + name + "! Welcome to Go programming."
}

func performOperation(x1 int, x2 int) int {
	result := x1 + x2 
	return result
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(customMessage("Alice"))
	fmt.Println(performOperation(10, 32))
}
