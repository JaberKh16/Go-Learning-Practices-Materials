package main

import "fmt"

// defining function
func add(a int, b int) int {
	return a + b
}


// example if having same type parameter
func mul(a, b int) int {
	return  a * b
}

// multiple returns => orders sequence matters here
func getLanguages()(string, string, string){
	return  "C++", "Golang", "JavaScript"
}

// function as first class citizen => used a variable
func firstClassFunc(fn func(a int, b int) int) int {
	return fn(1, 10)
}


func main(){

	// general use
	result := add(100, 32)
	fmt.Println(result)

	result2 := mul(10, 45)
	fmt.Println(result2)

	// multiple value
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)


	// first class citizen use case
	fn := func(a int, b int) int {
		return a + b
	}
	returnValue := firstClassFunc(fn)
	fmt.Println(returnValue)
}