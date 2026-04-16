/*
	Switch Case In Go
	===========================
	In Go, the `switch` statement is a powerful control structure that allows you to execute different blocks of code based on the value of an expression. It is often used as a more concise and readable alternative to multiple `if-else` statements.

	The syntax for a `switch` statement in Go is as follows:

	```go
	switch expression {
	case value1:
		// code to execute if expression equals value1
	case value2:
		// code to execute if expression equals value2
	default:
		// code to execute if expression does not match any case
	}
	```

	In the example below, we use a `switch` statement to determine the name of a month based on its corresponding number. If the month number does not match any of the cases, we print "Other month".

	Note: In Go, the `switch` statement does not require a `break` statement at the end of each case, as it automatically breaks after executing a case. However, if you want to continue executing the next case, you can use the `fallthrough` keyword.
*/

package main

import "fmt"

func main() {

	month := 20

	switch month {
	case 1:
		fmt.Println("January")
		// fallthrough // This will cause the execution to continue to the next case
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	default:
		fmt.Println("Other month")
	}
}