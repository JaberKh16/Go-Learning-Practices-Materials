package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Manufacture struct {
	Name    string `json:"name"`
	Flavour string `json:"flavour"`
}

type OSKernels struct {
	Name         string      `json:"name"`
	Version      string      `json:"version"`
	ReleasedYear string      `json:"released_year"`
	Manufacture  Manufacture `json:"manufacture"`
}

func main() {

	file, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var osListMap map[int] OSKernels

	err = json.Unmarshal(file, &osListMap)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Loaded Data:")
	for k, v := range osListMap {
		fmt.Printf("%d => %+v\n", k, v)
	}
}