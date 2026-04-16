package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	age := 20

	// typecasting using parsing
	ageStr := fmt.Sprintf("%d", age) // Convert int to string using Sprintf

	fmt.Println("Age (int):", age)
	fmt.Println("Age (string):", ageStr)
	fmt.Println("Type of age:", fmt.Sprintf("%T", age))
	fmt.Println("Type of ageStr:", fmt.Sprintf("%T", ageStr))


	// To convert a string back to an integer, you can use the strconv package
	ageParsed, err := strconv.Atoi(ageStr) // Convert string back to int using Atoi
	fmt.Println("Type of ageParsed, err:", fmt.Sprintf("%T", ageParsed), fmt.Sprintf("%T", err))
}
