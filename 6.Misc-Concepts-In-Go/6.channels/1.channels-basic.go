/*
	Channels Concepts In Go
	========================

	# What Are Channels?

	Channels are a way for goroutines to communicate with each other.
	They allow you to send and receive values between concurrent processes safely.

	Go follows the principle:
	"Do not communicate by sharing memory; share memory by communicating."

	# Channel Syntax

	Create a channel using make:

		ch := make(chan int)

	This creates a channel that can send and receive integers.

	# Sending and Receiving

	- Send data into a channel:
		ch <- 10

	- Receive data from a channel:
		value := <-ch

	# Blocking Behavior

	Channels are blocking by default:

	- Sending blocks until another goroutine receives
	- Receiving blocks until another goroutine sends

	This helps synchronize goroutines automatically.

	# Buffered Channels

	You can create channels with capacity:

		ch := make(chan int, 2)

	- Can hold 2 values without blocking
	- Useful when you want limited buffering

	# Closing Channels

	Channels can be closed using:

		close(ch)

	- Indicates no more values will be sent
	- Receivers can still read remaining values

	# Range Over Channel

	You can loop over a channel:

		for val := range ch {
			fmt.Println(val)
		}

	Loop ends automatically when channel is closed.

	# Directional Channels

	You can restrict channel usage:

		func sendData(ch chan<- int)   // send only
		func receiveData(ch <-chan int) // receive only

	Improves safety and clarity.

	# Select Statement

	Used to handle multiple channels:

		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case ch2 <- 10:
			fmt.Println("sent")
		default:
			fmt.Println("no activity")
		}

	# Use Cases of Channels

	1. Communication between goroutines
	2. Synchronization without locks
	3. Worker pools
	4. Pipelines (data processing stages)
	5. Handling async tasks

	# Key Idea

	Channels provide a safe and structured way for goroutines
	to communicate and synchronize without explicit locks.
*/

package main

import (
	"fmt"
	"time"
)

// Generic function (fixed)
func CaseDeadlock[T any](ch chan T, data T) T {
	// send data
	ch <- data

	// receive data
	result := <-ch
	return result
}

func goRoutineSetup(ch chan int) {
	go func() {
		ch <- 42 // send or set value
	}()
}

func processRandomValue(processes chan int){
	for process := range processes{
		fmt.Println(process)
		time.Sleep(time.Microsecond * 3000)
	} 
}

func main() {
	// create channel
	ch := make(chan int)

	// start goroutine that sends value
	goRoutineSetup(ch)

	// receive value
	value := <-ch

	fmt.Println("Received:", value)

	// Example of generic function usage
	ch2 := make(chan string)
	go func() {
		result := CaseDeadlock(ch2, "Deadlock")
		time.Sleep(time.Second * 3) // sleep for 3 secs
		fmt.Println("Generic Result:", result)
	}()

	// process random values
	go processRandomValue(ch)
	time.Sleep(time.Second * 2) // let goroutine process some values
}