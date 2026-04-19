package main

import "fmt"

func main() {
	// Define an anonymous function to calculate the square of a number
	square := func(x int) int {
		return x * x
	}

	// Call the anonymous function and print the result
	result := square(5)
	fmt.Println("The square of 5 is:", result)
}