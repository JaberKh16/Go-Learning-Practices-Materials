package main

import "fmt"

func main() {
	
	// variable declaration in conditionals
	if value := 20; value >= 15 {
		fmt.Println("value is greater than or equal to 15")
	} else if value := 10; value >= 10 {
		fmt.Println("value is between 10 and 14")
	} else if value >= 5 {
		fmt.Println("value is between 5 and 9")
	} else {
		fmt.Println("value is less than 5")
	}
}
