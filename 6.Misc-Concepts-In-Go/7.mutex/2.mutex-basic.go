/*
	MUTEX CONCEPTS IN GO
	====================

	# What is a Mutex?

	A Mutex (Mutual Exclusion) is a synchronization tool
	used to protect shared data from being accessed by
	multiple goroutines at the same time.

	It ensures that only ONE goroutine can access a
	critical section at a time.

	# Why Mutex is Needed?

	When multiple goroutines access and modify the same
	variable, it can cause a race condition:

	- Data becomes inconsistent
	- Results become unpredictable
	- Bugs are hard to trace

	Example problem:
	Two goroutines increment the same counter simultaneously
	can lead to wrong final values.

	# Basic Mutex Syntax

	Import sync package:

		import "sync"

	Create a mutex:

		var mu sync.Mutex

	# Lock and Unlock

	Mutex has two main methods:

		mu.Lock()   -> lock access (enter critical section)
		mu.Unlock() -> release access

	Only one goroutine can hold the lock at a time.

	# Critical Section

	The part of code where shared data is accessed:

		mu.Lock()
		// shared data access (critical section)
		mu.Unlock()

	This ensures safe access.

	# Safe Counter Example

	Using mutex to protect shared variable:

		var counter int
		var mu sync.Mutex

		func increment() {
			mu.Lock()
			counter++
			mu.Unlock()
		}

	# Using WaitGroup with Mutex

	WaitGroup ensures goroutines finish execution:

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()

		wg.Wait()

	# Race Condition vs Mutex

	Without Mutex:
	- Multiple goroutines modify data at same time
	- Leads to incorrect results

	With Mutex:
	- Only one goroutine modifies data at a time
	- Ensures correctness

	# Locking Best Practice

	Always unlock properly:

		mu.Lock()
		defer mu.Unlock()

	This avoids deadlocks if function exits early.

	# Deadlock (Important Concept)

	A deadlock happens when:

	- A goroutine locks but never unlocks
	- Other goroutines keep waiting forever

	Example mistake:
		mu.Lock()
		// forgot mu.Unlock() ❌

	# Types of Mutex

	1. sync.Mutex
		- Basic lock (only one goroutine allowed)

	2. sync.RWMutex
		- Multiple readers OR one writer
		- Better performance for read-heavy systems

	# RWMutex Idea

		RLock()   -> multiple readers allowed
		RUnlock() -> release read lock
		Lock()    -> write lock (exclusive)
		Unlock()  -> release write lock

	# Use Cases of Mutex

	1. Protect shared variables
	2. Protect maps (important in Go)
	3. File writing synchronization
	4. Shared counters / caches
	5. Backend APIs with shared state

	# Key Idea

	Mutex ensures safe access to shared memory
	by allowing only one goroutine at a time
	in the critical section.
*/

package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup){
	defer wg.Done()
	p.mu.Lock()
	p.views += 1
	p.mu.Unlock()
}

func (p *post) dec(wg *sync.WaitGroup){
	// safer access catches exception issue
	defer func ()  {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.views += 1

}

func main() {
	var wg sync.WaitGroup

	postViews := post{views: 4}

	for idx:=1; idx < 100; idx++{
		wg.Add(1)
		go postViews.inc(&wg)
	}

	wg.Wait()
	fmt.Println(postViews)
}