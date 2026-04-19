/*
	String In Go
	============
	In Go, a string is a sequence of characters enclosed in double quotes. Strings are immutable, meaning that once a string is created, it cannot be changed. You can create a string variable and assign it a value using the `var` keyword or the shorthand `:=` syntax.

	In the example below, we declare a string variable named `str` and assign it the value "Hello, Go!". We then print the value of `str` to the console.
*/

package main

import "fmt"

func main() {
	var str string = "Hello, Go!" // string must be enclosed in double quotes


	// You can also use shorthand syntax to declare and initialize a string variable
	// str := "Hello, Go!" // This is an alternative way to declare and initialize a string variable

	// using string literal
	strLiteral := `Hello, Go!` // This is another way to declare a string using backticks, which allows for multi-line strings
	fmt.Println(str)
	fmt.Println(strLiteral)


	// looping
	for index, unicode := range str {
		fmt.Println(index, unicode, string(unicode))
	}

}
