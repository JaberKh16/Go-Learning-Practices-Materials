package main

import (
	"fmt"
)

func printArr(arr []int) {
	// %v: prints the value in default format
	fmt.Printf("%v\n", arr)
	
	// %+v: prints with field names for structs
	// fmt.Printf("%+v\n", arr)
	
	// %#v: prints Go syntax representation
	// fmt.Printf("%#v\n", arr)
	
	// %T: prints the type of the value
	// fmt.Printf("%T\n", arr)
}

func performOperation(option string,  arr [] int )  {
	fmt.Println("Enter the option: ")
	fmt.Scan(&option) 
}

func main(){
	arr := [] int {1, 2, 10, 423, 565, 12}

	// get the length
	fmt.Println(len(arr)) // returns [1, 2, 10, 423, 565, 12]

	// print the array
	printArr(arr)


	// append method
	// append - single
	arr1 := append(arr, 100) //  [1, 2, 10, 423, 565, 12]
	printArr(arr1)

	// append - add elements
	arr2 := append(arr, 50, 60)
	printArr(arr2)

	// slice - get portion of array
	slice := arr[1:4]
	printArr(slice)

	// copy - duplicate array
	arr3 := make([]int, len(arr))
	copy(arr3, arr)
	printArr(arr3)

	// delete element at index
	arr4 := append(arr[:2], arr[3:]...)
	printArr(arr4)

	// insert at index
	index := 2
	arr5 := append(append(arr[:index], 999), arr[index:]...)
	printArr(arr5)



	// get user input
	var option int
	

	// switch expression {
	// case condition:
		
	// }
	
}