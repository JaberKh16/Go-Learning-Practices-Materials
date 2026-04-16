/*
Infered Type Variables in Go
============================
In Go, you can declare variables without explicitly specifying their type. The Go compiler will infer
the type based on the value assigned to the variable. This is known as "type inference".

In the example below, we declare variables using the short variable declaration syntax (:=),
which allows us to omit the type. The compiler will automatically determine the type of
each variable based on the assigned value.
*/
package main

import "fmt"

func main() {

	name := "Go Programming"
	age := 10
	salary := 50000.50
	isFun := true


	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Is Go fun?", isFun)
}