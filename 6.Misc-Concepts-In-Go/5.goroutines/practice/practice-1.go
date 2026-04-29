package main

import (
	"fmt"
	"math/rand"
	"time"
)

func braodCast(c chan int)  {
	for {
		c <- rand.Intn(1000)
	}
}


func main() {
	numbersStation := make(chan int)

	// execute broadcast in a separate thread
	go braodCast(numbersStation)

	for num := range numbersStation {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%d", num)
	}

}