package main

import "fmt"

// define a custom type
type CustomType interface{}


// generic function => via any
func printSlice[T any](items []T){
	for _, item := range items{
		fmt.Println(item)
	}
}

// generic function => via interface
func printInterfaceSlice[T interface{}](items []T){
	for _, item := range items{
		fmt.Println(item)
	}
}

// scoped generic function => via specified type => type | type
func printScopedIntefaceSlice[T int | string | bool](items []T){
	for _, item := range items {
		fmt.Println(item)
	}
}

func main(){
	printSlice([]int {1, 12, 32, 45})
	printSlice([]string {"C++", "JavaScript", "TypeScript"})
}