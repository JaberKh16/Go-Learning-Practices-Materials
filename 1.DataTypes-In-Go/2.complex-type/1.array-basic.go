package main

func main() {
	var arr [5]int // Declare an array of 5 integers

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
