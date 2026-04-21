package main

import (
	"fmt"
	"time"
)


func main(){

	// define struct type => a collection of elements called field where each field contains name and type
	person1 := struct{
		Name string
		Age int
		Gender string
		BornAt time.Time
		CreatedAt time.Time
	}{

		Name: "Mr. X",
		Age: 30,
		Gender: "male",
		CreatedAt: time.Now(),
	}
	
	fmt.Println(person1)
}