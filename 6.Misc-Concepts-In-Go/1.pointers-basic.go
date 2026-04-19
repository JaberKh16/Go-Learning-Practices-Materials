package main

import "fmt"

func changeByValue(value int){
	value = 10
	fmt.Println("Changed value: ", value)
}

func changeByReference(value *int){
	*value = 10 // derefence the pointer
	fmt.Println("Changed value: ", value)
}

func main(){
	value :=12
	changeByValue(value) // pass the value
	fmt.Println("Value is: ", value)
	fmt.Println("Address is: ", &value)


	changeByReference(&value) // pass hhe reference
	fmt.Println("Value is: ", value)
	fmt.Println("Address is: ", &value)
}