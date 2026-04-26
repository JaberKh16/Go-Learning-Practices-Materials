package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
	Gender string
}

func main(){
	person := Person {
		Name: "Mr. X",
		Age: 30,
		Gender: "male",
	}

	// convert to JSON Encoding(Marshalling)
	jsonFormattedData, _ := json.Marshal(person) // json.Marshal() returns a byte array
	fmt.Println(string(jsonFormattedData))

	// decode(Unmarshal)
	var decodedFormattedData Person
	err := json.Unmarshal(jsonFormattedData, &decodedFormattedData)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	fmt.Println("Decoded data is: ", decodedFormattedData)


}