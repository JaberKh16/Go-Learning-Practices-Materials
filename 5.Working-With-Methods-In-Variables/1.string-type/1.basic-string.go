package main

import "strings"


func main() {
	str := "Hello, World!"
	
	// Convert to uppercase
	upperStr := strings.ToUpper(str)
	println("Uppercase:", upperStr)

	// Convert to lowercase
	lowerStr := strings.ToLower(str)
	println("Lowercase:", lowerStr)

	// Check if the string contains a substring
	contains := strings.Contains(str, "World")
	println("Contains 'World':", contains)

	// Split the string into a slice
	splitStr := strings.Split(str, ", ")
	println("Split string:", splitStr[0], "and", splitStr[1])

	// Trim whitespace
	trimmedStr := strings.TrimSpace("   Hello, Go!   ")
	println("Trimmed string:", trimmedStr)

	// Replace a substring
	replacedStr := strings.Replace(str, "World", "Go", 1)
	println("Replaced string:", replacedStr)

	// Replace all occurrences of a substring
	replacedAllStr := strings.ReplaceAll(str, "o", "0")
	println("Replaced all 'o' with '0':", replacedAllStr)

	// Get the length of the string
	length := len(str)
	println("Length of string:", length)

	// Check if the string has a prefix
	hasPrefix := strings.HasPrefix(str, "Hello")
	println("Has prefix 'Hello':", hasPrefix)

	// Check if the string has a suffix
	hasSuffix := strings.HasSuffix(str, "!")
	println("Has suffix '!':", hasSuffix)
}