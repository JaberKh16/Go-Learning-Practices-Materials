package main

import (
	"fmt",
	"math"
)

func main() {
	x := 5

	result := square(x)
	fmt.Printf("The square of %d is %d\n", x, result)

	result = cube(x)
	fmt.Printf("The cube of %d is %d\n", x, result)

	// using math.Pow for square
	result = math.Pow(float64(x), 2)
	fmt.Printf("The square of %d is %.0f\n", x, result)

	// using math.Pow for cube
	x = 5.234234
	result = math.Ceil(float64(x))
	fmt.Printf("The cube of %d is %.0f\n", x, result)
}
