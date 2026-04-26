package main

import "fmt"

func main() {
	var a interface{} = "some" // define a interface having type string
	var b any = "some"         // define any having type string i.e interface {}

	c := a.(string) // if declared interface {} or any not type of string it will panic
	fmt.Println(c)

	// check on panic situation
	if c, ok := a.(string); ok {
		fmt.Println(c)
	}
}
