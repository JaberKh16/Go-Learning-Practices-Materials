package main

import "fmt"


func generateConditionalAccountNumber(accountNumberChannel chan int) {
	var accountNumber int 
	accountNumber = 3000001

	for  {
		// close the channel after 5 numbers are requested
		if accountNumber > 3000001 {
			close(accountNumberChannel)
			return
		}
		accountNumberChannel <- accountNumber
		accountNumber +=1
	}
}

func generateAccountNumber(accountNumberChannel chan int) {
	// internal variable to store last generated account number
	var accountNumber int
	accountNumber = 3000001

	for {
		accountNumberChannel <- accountNumber
		accountNumber += 1
	}
}

func main() {
	accountNumberChannel := make(chan int)

	// start the goroutine that generates account numbers
	go generateAccountNumber(accountNumberChannel)

	fmt.Printf("SMITH: %d\n", <-accountNumberChannel)
	fmt.Printf("SINGH: %d\n", <-accountNumberChannel)
	fmt.Printf("JONES: %d\n", <-accountNumberChannel)
	fmt.Printf("LOPEZ: %d\n", <-accountNumberChannel)
	fmt.Printf("CLARK: %d\n", <-accountNumberChannel)
}
