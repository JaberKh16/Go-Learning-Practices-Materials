package main

import (
	"fmt"
	"time"
)

func main(){

	// define struct type => a collection of elements called field where each field contains name and type
	type Person struct{
		Name string
		Age int
		Gender string
		BornAt time.Time
	}

	fmt.Println(Person)
}