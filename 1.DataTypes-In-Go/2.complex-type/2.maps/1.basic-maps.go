package main

import "fmt"

func main(){
	// maps -> hash, object, dict
	maps := make(map[string]string)

	// assign value
	maps["name"] = "Mr. X"
	maps["designation"] = "Sr. Software Engineer"

	// print the map
	fmt.Println(maps)

	// get the length
	fmt.Println(len(maps))
}
