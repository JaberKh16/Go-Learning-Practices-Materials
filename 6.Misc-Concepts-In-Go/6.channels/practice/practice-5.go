// Practical Example: Multiple workers + select
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker", id, "processing", job)
		time.Sleep(time.Millisecond * 500)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Fan-out: start multiple workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for i := 0; i < 5; i++ {
		fmt.Println("Result:", <-results)
	}
}