package main

import (
	"fmt"
	"sort"
)

func main(){
	
	// define maps
	actor := map[string]int{
		"Paltrow": 43,
		"Cruise": 53,
		"Redford": 79,
		"Diaz": 43,
		"Kilmer": 56,
		"Pacino": 75,
		"Ryder": 44,
	}

	// store the keys in slice
	var sortedKeyList [] string
	var sortedValueList [] int
	for key, value := range actor {
		sortedKeyList = append(sortedKeyList, key)
		sortedValueList = append(sortedValueList, value)
		// fmt.Println(key, sortedKeyList, value)
	}

	// sort the slice
	sort.Strings(sortedKeyList)
	// sort.Strings(sortedValueList)
	fmt.Println(sortedKeyList)
	fmt.Println(sortedValueList)




}