package main

import (
	"fmt"
	"math"
)

func main() {
	x := 5

	// square using math.Pow
	result := math.Pow(float64(x), 2)
	fmt.Printf("The square of %d is %.0f\n", x, result)

	// cube using math.Pow
	cube := math.Pow(float64(x), 3)
	fmt.Printf("The cube of %d is %.0f\n", x, cube)

	// example of Ceil (rounding up a float)
	y := 5.234234
	ceilVal := math.Ceil(y)
	fmt.Printf("Ceil of %.6f is %.0f\n", y, ceilVal)
}