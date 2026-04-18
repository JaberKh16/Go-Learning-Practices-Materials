package main

import "fmt"

func main() {
	
	arr := [5]int{1, 2, 3, 4, 5} // Initialize an array with values

	// slice are basically the dynamic array in go
	// unintialized slice value is nil
	var slicedArr []int
	fmt.Println(slicedArr == nil) // returns true
	fmt.Println(len(slicedArr)) // returns 0


	// Slicing the array to create a slice
	slice := arr[1:4] // This will create a slice containing elements from index 1 to 3 (4 is exclusive)

	// Print the original array and the sliced slice
	fmt.Println("Original Array:", arr)	
	fmt.Println("Sliced Slice:", slice)
	

	// mutating the slice will also mutate the original array since they share the same underlying array
	// mutation makes the chnges in both arr and slice since they share the same underlying array
	slice[0] = 20 // This will change arr[1] to 20


	println("==================After Mutation===========================")

	fmt.Println("Original Array after mutation:", arr) // arr will reflect the change made through the slice
	fmt.Println("Sliced Slice:", slice)


	println("==================After Mutation Check Array Properties===========================")

	// pointer -> both arr and slice point to the same underlying array in memory
	// length -> arr has a length of 5, while slice has a length of 3 (from index 1 to 3)
	// capacity -> the capacity of the slice is determined by the length of the original array from the starting index of the slice to the end of the array, so in this case, the capacity of the slice is 4 (from index 1 to 4)

	// Check the length and capacity of the original array and the sliced slice
	fmt.Printf("Length of Original Array: %d\n", len(arr)) // Length of the original array
	fmt.Printf("Capacity of Sliced Slice: %d\n", cap(slice)) // Capacity of the sliced slice

}
