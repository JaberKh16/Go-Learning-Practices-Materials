package main

import "fmt"

type OSKernels struct {
	Name string
	Version string 
	ReleasedYear string 
}

func main() {

	// define map
	osListMap := make(map[int] OSKernels)

	// fmt.Println(osListMap)

	// insert 
	// osListMap["pack1"] = OSKernels{Name: "Windows", Version: "v10", ReleasedYear: "2010"}

	for i:=0; i < 3; i++ {
		if i == 0 {
			osListMap[i+1] = OSKernels{Name: "Windows", Version: "v10", ReleasedYear: "2010"}
		} else if i == 1 {
			osListMap[i+1] = OSKernels{Name: "Linux", Version: "kernel-16.21.2", ReleasedYear: "2012"}
		} else {
			osListMap[i+1] = OSKernels{Name: "MacOS", Version: "mckernel-8.2.12", ReleasedYear: "2009"}
		}
	}

	fmt.Println(osListMap)
	
}