package main

import "fmt"

func main(){

	arr := make([] int, 2, 5) // syntax: make(type, size, cap)


	// append items => appends at end
	nums := append(arr, 100)
	nums = append(nums, 200)
	nums = append(nums, 300)
	nums = append(nums, 400)
	nums = append(nums, 500)
	nums = append(nums, 600)

	// list properties
	fmt.Println("Array: ", nums)
	fmt.Println("Length: ",len(nums))
	fmt.Println("Capacity: ",cap(nums))

	// perform copy
	nums2 := make([]int, len(nums), 2)
	copy(nums2, nums)
	fmt.Println(nums2)
}
