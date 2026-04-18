package main

import "fmt"

func main() {
	var arr [5]int // Declare an array of 5 integers

	// print the array
	fmt.Println("Initial Array: ",arr) // [0, 0, 0, 0, 0] though type is int


	// Initialize the array with values
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	// Print the array
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
}
