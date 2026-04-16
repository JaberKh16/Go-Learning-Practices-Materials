/*
	Typecasting Variable In Go
	===========================
	In Go, you can perform typecasting to convert a variable from one data type to another. This is done using the syntax `Type(variable)`, where `Type` is the desired data type and `variable` is the value you want to convert.

	In the example below, we declare a variable `age` of type `int` and assign it the value `10`. We then use typecasting to convert `age` to a `float64` and store it in a new variable called `ageFloat`. Finally, we print both the original integer value and the converted float value.
*/

package main

import "fmt"

func main() {
	var age int = 10
	var ageFloat float64 = float64(age) // Typecasting 'age' from int to float64

	fmt.Println("Age (int):", age)
	fmt.Println("Age (float64):", ageFloat)
	fmt.Println("Type of age:", fmt.Sprintf("%T", age))
	fmt.Println("Type of ageFloat:", fmt.Sprintf("%T", ageFloat))
}

