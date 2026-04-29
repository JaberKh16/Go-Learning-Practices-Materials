package main

import (
	"fmt"
	"sync"
)

type AccountGenerator struct {
	mu sync.Mutex
	current int 
}

func NewAccountGenerator(start int) *AccountGenerator {
	return  &AccountGenerator{
		current: start,
	}
}


func (g *AccountGenerator) Next() int  {
	g.mu.Lock()
	defer g.mu.Unlock()

	acc := g.current
	g.current++
	return  acc
}


func main()  {
	gen := NewAccountGenerator(30001)
	var wg sync.WaitGroup

	// simulate concurrent requests
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			fmt.Println("Account: ", gen.Next())
		}()
	}
	wg.Wait()
}