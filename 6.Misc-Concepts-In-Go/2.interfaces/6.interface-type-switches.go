package main

import "fmt"

func process[T any] (value T){
	// gets the any type , type
	switch v := any(value).(type) {
	case int:
		fmt.Printf("type: %s", v)
	case float64:
		fmt.Println("float: ", v)
	case string:
		fmt.Println("string: ", v)
	case byte:
		fmt.Println("string: ", v)
	default: 
		fmt.Println("unknown type")
	}  
}

func main(){
	val := 3
	process(val)
}