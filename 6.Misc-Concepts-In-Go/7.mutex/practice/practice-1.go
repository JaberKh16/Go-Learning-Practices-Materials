package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mu.Lock()   // 🔒 lock critical section
		counter++
		mu.Unlock() // 🔓 unlock
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}