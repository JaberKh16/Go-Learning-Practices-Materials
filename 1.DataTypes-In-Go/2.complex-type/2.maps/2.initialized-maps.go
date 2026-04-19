package main

import (
	"fmt"
)

func main(){
	// maps -> hash, object, dict
	mapsData := map[string]int{
		"port":  5000,
		"connectionLimit": 10,
		"status": 200,
	}

	// print the map
	fmt.Print(mapsData["port"]) // if found key then return the value, otherwise return empty value 


	// check 
	value, exists := mapsData["port"] // returns value if exits in 'value' and exits returns bool if exists
	if exists {
		fmt.Println("Exists:", value)
	} else {
		fmt.Println("Not found")
	}


	// delete key
	delete(mapsData, connectionLimit)

	// clear the map
	clear(mapsData)


}
