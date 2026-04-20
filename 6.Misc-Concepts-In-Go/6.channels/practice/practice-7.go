// Rest API Based Examples
package main

import (
	"fmt"
	"time"
)

// Simulated API calls
func fetchUser(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "User: John"
}

func fetchOrders(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Orders: 5"
}

func fetchRecommendations(ch chan<- string) {
	time.Sleep(3 * time.Second)
	ch <- "Recommendations: 10 items"
}

func main() {
	userCh := make(chan string)
	orderCh := make(chan string)
	recCh := make(chan string)

	// Fan-out: call APIs concurrently
	go fetchUser(userCh)
	go fetchOrders(orderCh)
	go fetchRecommendations(recCh)

	timeout := time.After(2 * time.Second)

	results := make(map[string]string)

	// Fan-in using select
	for i := 0; i < 3; i++ {
		select {
		case user := <-userCh:
			results["user"] = user

		case orders := <-orderCh:
			results["orders"] = orders

		case rec := <-recCh:
			results["recommendations"] = rec

		case <-timeout:
			fmt.Println("Timeout reached!")
			break
		}
	}

	// Final aggregated response
	fmt.Println("Final Response:")
	for k, v := range results {
		fmt.Println(k, "=>", v)
	}
}