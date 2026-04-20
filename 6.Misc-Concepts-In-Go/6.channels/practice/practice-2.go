package main

import (
	"fmt"
	"math/rand"
)

// receive channel value function
func performOperation(result chan int, num1 int, num2 int){
	operationResult := num1 + num2
	result <- operationResult // pass the data to the channel
}

func main() {
	// create channel
	result := make(chan int)

	// goroutine
	go performOperation(result, rand.Intn(100), rand.Intn(50))

	// receive
	receiveResult := <-result
	fmt.Println(receiveResult)
	
}