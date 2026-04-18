package main

import "fmt"

func main(){
	// maps -> hash, object, dict
	maps := map[string]int{
		"port":  5000
		"connectionLimit": 10
		"status": 200
	}

	// print the map
	fmt.Print(maps["port"]) // if found key then return the value, otherwise return empty value 


	// check 
	value, exists := maps["port"]
	if exists {
		fmt.Println("Exists:", value)
	} else {
		fmt.Println("Not found")
	}


	// delete the data
	delete(maps, connectionLimit)

}
