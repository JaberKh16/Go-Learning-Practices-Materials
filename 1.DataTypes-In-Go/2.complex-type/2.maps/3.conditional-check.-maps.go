package main

import (
	"fmt"
	"maps"
)

func main(){
	// maps -> hash, object, dict
	mapsData := map[string]int{
		"port":  5000,
		"connectionLimit": 10,
		"status": 200,
	}


	// check 
	value, exists := mapsData["port"] // returns value if exits in 'value' and exits returns bool if exists
	if exists {
		fmt.Println("Exists:", value)
	} else {
		fmt.Println("Not found")
	}


	// another maps
	locationMaps1 := map[string]float64{
		"latitude":  40.7128,
		"longitude": -74.0060,
		"altitude":  10.5,
	}

	locationMaps2 := map[string]float64{
		"latitude":  40.7128,
		"longitude": -74.0060,
		"altitude":  10.5,
	}


	// check equality
	// fmt.Println(locationMaps1 == locationMaps2) // exception direct check not available
	fmt.Println(maps.Equal(locationMaps1, locationMaps2))


}
