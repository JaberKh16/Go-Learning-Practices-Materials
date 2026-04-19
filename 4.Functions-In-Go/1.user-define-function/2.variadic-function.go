/*
	Variadic Function In GO
	========================
	Function that takes any values as parameters are variadic functions.
	Syntax:
		func funcName(p1 ...type) type {
		}
*/

package main

import "fmt"

// use case: 1 using spread operator
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main(){
	result := sum(1, 313, 1313, 46, 24, 242)
	fmt.Println(result)


	// using slice
	arr := [] int {10, 12, 23}
	result2 := sum(arr...)
	fmt.Println(result2)
}
