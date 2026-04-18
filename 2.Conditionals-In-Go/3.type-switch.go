package main

import (
	"fmt"
	"reflect"
)

func main(){

	typeSwitch := func(i interface {}){
		switch t := i.(type) {
		case int: 
			fmt.Println("integer: ", t, reflect.TypeOf(t))
		case float64: 
			fmt.Println("float: ", t, reflect.TypeOf(t))
		case string: 
			fmt.Println("string: ", t, reflect.TypeOf(t))
		case bool: 
			fmt.Println("boolean: ", t, reflect.TypeOf(t))
		default: 
			fmt.Println("other type: ", t, reflect.TypeOf(t))
		}
	}

	// calling the function
	typeSwitch(42)
	typeSwitch("soemthing")
	typeSwitch(42.242)
}
