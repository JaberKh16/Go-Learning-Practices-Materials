package main

import "fmt"

func main() {
	arr := [3][2]int{ // Declare a nested array (array of arrays)
		{1, 2}, // First inner array
		{3, 4}, // Second inner array
		{5, 6}, // Third inner array
	}

	// Print the nested array
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d ", arr[i][j]) // Access elements using arr[i][j]
		}
		fmt.Println() // New line after each inner array
	}
}
