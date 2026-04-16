package main

import "fmt"

func main() {
	var runeValue rune = 'A' // A rune is a Unicode code point, and it can be represented as a character literal
	fmt.Printf("The rune value is: %c\n", runeValue) // Print the character representation of the rune
	fmt.Printf("The Unicode code point of the rune is: %U\n", runeValue) // Print the Unicode code point of the rune
}
