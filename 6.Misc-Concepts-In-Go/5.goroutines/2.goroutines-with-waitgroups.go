/*
	Goroutines Concepts In Go
	=========================


	# What Are Goroutines?

	A goroutine is a lightweight thread managed by the Go runtime.
	It allows functions to run concurrently with other functions.

	Goroutines are one of the key features of Go for concurrency.

	# How to Start a Goroutine

	You start a goroutine by using the "go" keyword before a function call:

		go functionName()

	This runs the function concurrently in the background.

	# Main Goroutine

	Every Go program starts with a main goroutine.
	If the main function finishes, all other goroutines are stopped immediately.

	# Concurrency vs Parallelism

	- Concurrency: multiple tasks are in progress at the same time
	- Parallelism: multiple tasks run at the same time on multiple CPU cores

	Goroutines enable concurrency and can achieve parallelism when needed.

	# Goroutine Characteristics

	- Very lightweight (starts with small stack ~2KB)
	- Managed by Go runtime scheduler
	- Thousands or even millions can run efficiently

	# Synchronization Problem

	Goroutines run independently, so timing is not guaranteed.
	This may lead to race conditions if multiple goroutines access shared data.

	Solution: use synchronization tools like:
	- channels
	- mutexes (sync.Mutex)
	- WaitGroups

	# Use Cases of Goroutines

	1. Handling multiple web requests
	2. Background tasks (logging, cleanup jobs)
	3. Parallel data processing
	4. Real-time systems (chat apps, streaming)
	5. API calls in parallel

	# Key Idea

	Goroutines allow functions to run concurrently in a lightweight and efficient way.
	They are the foundation of concurrency in Go.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)



func taskProces(limit int, w *sync.WaitGroup) {
	defer w.Done()
	for i := 1; i <= limit; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond)
	}
}


func exampleMain() {

	// define a variable
	var limit int = 10

	// defining waitgroups
	var wg sync.WaitGroup

	// Add one goroutine to the WaitGroup
	wg.Add(1)

	// Start goroutine
	go taskProces(limit, &wg)

	// Main goroutine work
	fmt.Println("Main function running")

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Main function finished")
}