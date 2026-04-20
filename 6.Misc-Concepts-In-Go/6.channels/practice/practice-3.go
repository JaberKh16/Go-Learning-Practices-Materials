package practice

import "fmt"
package main

import (
	"fmt"
	"math/rand"
)

// receive channel value function
func processTask(job chan bool){
	defer func (){
		job <- true
	}()
	fmt.Println("processing.....")
}

func main() {
	// create channel
	result := make(chan bool)

	// goroutine
	go processTask(result)

	// receive
	<-result // blocking
	
}