package main

import "fmt"

func main() {
	
	msg := "Something is wrong"

	// print the first character of the string
	fmt.Println("First character of the string:", msg[0]) // prints the uncide value of the first character 'S' which is 83

	// print the first character of the string as a string
	fmt.Println("First character of the string as a string:", string(msg[0])) // prints 'S' as a string

	// You can access individual characters in a string using indexing
}
