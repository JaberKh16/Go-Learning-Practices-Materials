// Advanced: Graceful shutdown with select

package main

import (
	"fmt"
)

func producer(name string, out chan<- string) {
	for i := 1; i <= 3; i++ {
		out <- fmt.Sprintf("%s: %d", name, i)
	}
}

// Fan-in function
func fanIn(ch1, ch2 <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for {
			select {
			case v := <-ch1:
				out <- v
			case v := <-ch2:
				out <- v
			}
		}
	}()

	return out
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go producer("A", ch1)
	go producer("B", ch2)

	merged := fanIn(ch1, ch2)

	for i := 0; i < 6; i++ {
		fmt.Println(<-merged)
	}
}