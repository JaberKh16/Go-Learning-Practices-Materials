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

}

func main(){
	arr := [] int {1, 2, 10, 423, 565, 12}

	// get the length
	fmt.Println(len(arr)) // returns [1, 2, 10, 423, 565, 12]

	// print the array
	printArr(arr)


	// append method
	arr1 := append(arr, 100) //  [1, 2, 10, 423, 565, 12]
	printArr(arr1)



	// switch expression {
	// case condition:
		
	// }
	
}