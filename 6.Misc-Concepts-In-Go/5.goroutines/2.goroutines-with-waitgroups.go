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

	========================================================
	⚠️ Concurrency Problems
	========================================================

	# 1. Race Condition

	A race condition happens when multiple goroutines access and modify
	shared data at the same time without synchronization.

	❌ Problem Example:

		var counter int

		func increment() {
			counter++ // unsafe
		}

		func main() {
			for i := 0; i < 1000; i++ {
				go increment()
			}
		}

		Result is unpredictable because multiple goroutines write at once.

	✅ Solution:
	- Use sync.Mutex
	- Use channels

	# 2. Deadlock

	A deadlock happens when goroutines are waiting on each other forever.

	❌ Problem Example:

		func main() {
			ch := make(chan int)

			ch <- 10 // no receiver → deadlock
		}

	This causes:

		fatal error: all goroutines are asleep - deadlock!

	Another example (circular wait):

		Goroutine A waits for B
		Goroutine B waits for A

	✅ Solution:
	- Ensure sender/receiver pairs exist
	- Avoid circular dependencies

	# 3. Livelock

	Livelock happens when goroutines keep changing state in response
	to each other but make no progress.

	❌ Example (conceptual):

		Two goroutines keep yielding to each other:

		for {
			select {
			case <-other:
				// give way
			default:
				// try again
			}
		}

	They are active but never complete useful work.

	✅ Solution:
	- Add delays/backoff (time.Sleep)
	- Reduce over-coordination

	# 4. Starvation

	Starvation happens when a goroutine never gets CPU time or resources
	because others dominate execution.

	❌ Example:

		func main() {
			go func() {
				for {
					// busy loop (never yields)
				}
			}()

			go func() {
				println("may never run")
			}()

			select {}
		}

	Second goroutine may not get enough execution time.

	✅ Solution:
	- Use runtime.Gosched()
	- Avoid tight infinite loops
	- Use proper scheduling (channels, waits)

	========================================================
	# Synchronization Tools

	Goroutines run independently, so timing is not guaranteed.

	To control execution:
	- channels
	- mutexes (sync.Mutex)
	- WaitGroups (sync.WaitGroup)

	========================================================
	# Use Cases of Goroutines

	1. Handling multiple web requests
	2. Background tasks (logging, cleanup jobs)
	3. Parallel data processing
	4. Real-time systems (chat apps, streaming)
	5. API calls in parallel

	========================================================
	# Key Idea

	Goroutines allow functions to run concurrently in a lightweight
	and efficient way.

	But improper use can lead to:
	- race conditions
	- deadlocks
	- livelocks
	- starvation

	Understanding these is critical for writing safe concurrent programs.
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
	now := time.Now()

	// defining waitgroups
	var wg sync.WaitGroup

	// Add one goroutine to the WaitGroup
	wg.Add(1)

	// Start goroutine
	go taskProces(limit, &wg)

	// Main goroutine work
	fmt.Println("Main function running")

	// elapsed
	fmt.Println("Time Elapased: ",time.Since(now))


	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Main function finished")
}