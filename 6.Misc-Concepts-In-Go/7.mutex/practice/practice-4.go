/* Atomic Counter */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicAccountGenerator struct {
	current int64
}

func NewAtomicGenerator(start int64) *AtomicAccountGenerator {
	return &AtomicAccountGenerator{
		current: start - 1,
	}
}

func (g *AtomicAccountGenerator) Next() int64 {
	return atomic.AddInt64(&g.current, 1)
}

func main()  {
	gen := NewAccountGenerator(30001)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Account: ", gen.Next())
		}()
	}
	wg.Wait()
}
