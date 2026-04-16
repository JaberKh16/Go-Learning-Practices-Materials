package main

import (
	"fmt"
	"reflect"
)

func main() {
	
	var x int = 42
	f := float64(x)

	fmt.Printf("x is of type: %s\n", reflect.TypeOf(x))
	// Output: x is of type: int

	fmt.Printf("f is of type: %s\n", reflect.TypeOf(f))
	// Output: f is of type: float64
}
