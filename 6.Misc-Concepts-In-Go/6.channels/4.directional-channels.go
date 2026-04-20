package main

import (
	"fmt"
	"time"
)

// Worker: receives jobs and sends results
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs { // receive-only
		fmt.Println("Worker", id, "processing job", job)
		time.Sleep(time.Second) // simulate work
		results <- job * 2      // send-only
	}
}

// Producer: sends jobs
func produceJobs(jobs chan<- int) {
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs) // important: signal no more jobs
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	go produceJobs(jobs)

	// Collect results
	for i := 1; i <= 5; i++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}