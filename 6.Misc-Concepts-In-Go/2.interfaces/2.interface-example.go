package main

import "fmt"

func getInfo(t interface {}) {
	fmt.Printf("Type: %T, Value: %v", t, t)
}

func main () {
	// define an empty interface {}
	mystreyBox := interface {} (10)

	// get the info
	getInfo(mystreyBox)


	// converting => type assertion
	interfaceVal, ok := mystreyBox.(int)
	if ok {
		fmt.Println("Data: ", interfaceVal)
	} else {
		fmt.Println("value conversion not happend")
	}

}