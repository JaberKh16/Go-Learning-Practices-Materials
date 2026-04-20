package main

import (
	"fmt"
	"math/rand"
	"time"
)

// send channel value function
func processRandomValue(processes chan int){
	for process := range processes{
		fmt.Println(process)
		time.Sleep(time.Microsecond * 3000)
	} 
}

// receive channel value function
func sum(result chan int, num1 int, num2 int){
	operationResult := num1 + num2
	result <- operationResult // pass the data to the channel
}

func main() {
	// create channel
	process := make(chan int)

	// process random values
	for {
		process <- rand.Intn(100) // random value from 0 to 100
	}
	
}