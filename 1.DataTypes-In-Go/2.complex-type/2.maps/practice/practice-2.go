package main

import (
	"encoding/json"
	"fmt"
)

type Manufacture struct {
	Name string
	Flavour string 
}

type OSKernels struct {
	Name string
	Version string 
	ReleasedYear string
	Manufacture 
}

func main() {

	// define map
	osListMap := make(map[int] OSKernels)

	// fmt.Println(osListMap)

	// insert 
	// osListMap["pack1"] = OSKernels{Name: "Windows", Version: "v10", ReleasedYear: "2010"}

	for i:=0; i < 3; i++ {
		if i == 0 {
			osListMap[i+1] = OSKernels{Name: "Windows", Version: "v10", ReleasedYear: "2010", Manufacture: Manufacture { Name: "Microsoft", Flavour: "Microsoft"}}
		} else if i == 1 {
			osListMap[i+1] = OSKernels{Name: "Linux", Version: "kernel-16.21.2", ReleasedYear: "2012", Manufacture: Manufacture { Name: "Red Hat Linux", Flavour: "Fedora"}}
		} else {
			osListMap[i+1] = OSKernels{Name: "MacOS", Version: "mckernel-8.2.12", ReleasedYear: "2009", Manufacture: Manufacture { Name: "Mactintosh", Flavour: "Mactintosh"}}
		}
	}

	// convert to json
	jsonData, err := json.MarshalIndent(osListMap, "", " ")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(osListMap)
	fmt.Println(string(jsonData))
	
}