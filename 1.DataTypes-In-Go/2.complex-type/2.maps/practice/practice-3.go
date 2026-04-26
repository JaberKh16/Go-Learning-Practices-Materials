package main

import (
	"encoding/json"
	"fmt"
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

	var n int
	fmt.Print("How many OS entries? ")
	fmt.Scanln(&n)

	osListMap := make(map[int] OSKernels)

	for i := 1; i <= n; i++ {
		var os OSKernels

		fmt.Printf("\n--- OS %d ---\n", i)

		fmt.Print("Name: ")
		fmt.Scanln(&os.Name)

		fmt.Print("Version: ")
		fmt.Scanln(&os.Version)

		fmt.Print("Released Year: ")
		fmt.Scanln(&os.ReleasedYear)

		fmt.Print("Manufacturer Name: ")
		fmt.Scanln(&os.Manufacture.Name)

		fmt.Print("Flavour: ")
		fmt.Scanln(&os.Manufacture.Flavour)

		osListMap[i] = os
	}

	jsonData, _ := json.MarshalIndent(osListMap, "", "  ")
	fmt.Println("\nGenerated JSON:\n", string(jsonData))
}