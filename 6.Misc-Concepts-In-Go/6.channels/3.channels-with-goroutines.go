package main

import (
	"fmt"
	"time"
)

func worker(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Work done"
}

func main() {
	ch := make(chan string)

	go worker(ch)

	fmt.Println("Waiting...")
	result := <-ch

	fmt.Println(result)
}