package main

import (
	"fmt"
	"time"
)

/// define struct type => a collection of elements called field where each field contains name and type
type Person struct{
	Name string
	Age int
	Gender string
	BornAt time.Time
	CreatedAt time.Time
}



func main(){

	// make an instance
	person1 := Person{
		Name: "Mr. X",
		Age: 30,
		Gender: "male",
		CreatedAt: time.Now(),
	}
	
	fmt.Println(person1)
}