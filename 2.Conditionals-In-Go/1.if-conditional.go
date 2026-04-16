/*
Conditionals In Go
===========================
In Go, you can use conditional statements to control the flow of your program based on certain conditions. The most common conditional statement is the `if` statement, which allows you to execute a block of code only if a specified condition is true.

The syntax for an `if` statement in Go is as follows:

```go

	if condition {
		// code to execute if condition is true
	} else {

		// code to execute if condition is false
	}

```

In the example below, we check if a person's age is greater than or equal to 18. If the condition is true, we print "You are an adult." Otherwise, we print "You are a minor."
*/

package main

import "fmt"

func main() {

	age := uint(20)

	if(age > 40) {
		fmt.Println("You are a senior.")
	} else if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}
}
