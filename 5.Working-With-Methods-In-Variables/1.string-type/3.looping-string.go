package main

func main() {
	
	str := "Hello, Go!"

	// traditional for loop to iterate over the string using index
	for idx := 0; idx < len(str); idx++ {
		char := str[idx] // got the unicode value of the character at index idx
		println("Index", idx, "Unicode Value:", char, "Character`:", string(char))
	}


	// Looping through each character in the string using a for loop
	// for idx, char := range str {
	// 	fmt.Printf("Character at index %d: %c\n", idx, char)
	// }


	// ignoring the index and just looping through the characters in the string
	// for _, char := range str {
	// 	fmt.Printf("Character: %c\n", char)
	// }

}
