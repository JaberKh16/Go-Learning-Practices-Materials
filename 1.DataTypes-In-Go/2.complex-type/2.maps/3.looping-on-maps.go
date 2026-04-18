package main

import "fmt"

func main(){
	// maps -> hash, object, dict
	maps := map[string]int{
		"port":  5000,
		"connectionLimit": 10,
		"status": 200,
	}

	// looping the data
	for key, value := range maps {
		fmt.Println(key, value)
	}
	
}
