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
	fmt.Println(str)
}
