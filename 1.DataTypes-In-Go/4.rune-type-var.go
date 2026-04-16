/*
    Rune Type in Go
	=========================
	In Go, a `rune` is a data type that represents a Unicode code point. It is an alias for the `int32` type and can hold any valid Unicode character. Runes are typically used to represent characters in strings, and they can be represented using character literals (enclosed in single quotes).

	In the example below, we declare a variable named `runeValue` of type `rune` and assign it the character 'A'. We then print both the character representation and the Unicode code point of the rune using formatted output.
*/

package main

import "fmt"

func main() {
	var runeValue rune = 'A' // A rune is a Unicode code point, and it can be represented as a character literal
	fmt.Printf("The rune value is: %c\n", runeValue) // Print the character representation of the rune
	fmt.Printf("The Unicode code point of the rune is: %U\n", runeValue) // Print the Unicode code point of the rune
}
