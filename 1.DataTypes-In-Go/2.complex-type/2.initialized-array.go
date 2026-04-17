package main

import "fmt"

func main() {
	
	arr := [5]int{1, 2, 3, 4, 5} // Initialize an array with values

	// Print the array
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
