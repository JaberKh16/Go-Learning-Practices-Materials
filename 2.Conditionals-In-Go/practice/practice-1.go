package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var randomNum int = rand.Intn(100)

	if randomNum > 50 && randomNum%2 == 0 {
		fmt.Println("It's closer to 100 and even")
	} else if randomNum < 50 {
		fmt.Println("It's closer to 0")
	} else if randomNum == 50 {
		fmt.Println("It's 50!")
	}

	fmt.Printf("Generated number: %v\n", randomNum)
}