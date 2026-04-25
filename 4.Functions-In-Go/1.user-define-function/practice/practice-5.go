package main

import (
	"fmt"
	"os"
	"strconv"
)

func oddOrEven(value string) string {
	num, _ := strconv.Atoi(value)
	if num % 2 > 0 {
		return "even"
	} else {
		return  "odd"
	}
}

func main() {
	// fmt.Printf("%s is %s\n", os.Args[1], oddOrEven(os.Args[1]))
	// fmt.Printf("%s is %s\n", os.Args[2], oddOrEven(os.Args[2]))
	// fmt.Printf("%s is %s\n", os.Args[3], oddOrEven(os.Args[3]))
	// fmt.Printf("%s is %s\n", os.Args[4], oddOrEven(os.Args[4]))

	for i:=1; i < len(os.Args); i++ {
		fmt.Printf("%s is %s\n", os.Args[i], oddOrEven(os.Args[i]))
	}
}