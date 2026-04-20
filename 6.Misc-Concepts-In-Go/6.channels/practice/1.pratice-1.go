package main

import (
	"fmt"
	"math/rand"
	"time"
)



func processRandomValue(processes chan int){
	for process := range processes{
		fmt.Println(process)
		time.Sleep(time.Microsecond * 3000)
	} 
}

func main() {
	// create channel
	process := make(chan int)


	// process random values
	for {
		process <- rand.Intn(100) // random value from 0 to 100
	}
}