package main

import (
	"fmt"
	"sync"
	"time"
)


// global variables
var (
	wg sync.WaitGroup
	results []string
	dbData = []string{"test1", "test2", "test3", "test4", "test5"}

)

func performOperation(i string) {
	var delay float32
	// simulate some work
	delay = 1.0
	// add delay
	time.Sleep(time.Duration(delay) * time.Second)
	// print each operation
	fmt.Printf("Operation: %s\n", i)
	// add to results
	results = append(results, i)
	wg.Done()
}

func main() {

	t0 := time.Now()
	for _, data := range dbData {
		wg.Add(1)
		go performOperation(data)
	}
	wg.Wait()
	fmt.Println("Results:", results)
	fmt.Println("Time taken:", time.Since(t0))
}
