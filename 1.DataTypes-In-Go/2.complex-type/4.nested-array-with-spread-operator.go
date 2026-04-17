package main

import "fmt"

func main() {
	// arr := [][]float32{
	// 	{1.0, 2.0},
	// 	{3.0, 4.0},
	// 	{5.0, 6.0},
	// }


	// using spread operator to let the compiler determine the length of the outer array
	arr := [...][2]float32{
		{1.0, 2.0},
		{3.0, 4.0},
		{5.0, 6.0},
	}
	
	for _, innerArr :=range arr {
		for _, value :=range innerArr {
			fmt.Printf("%.1f ", value)
		}
		fmt.Println()
	}
}