package main

import "fmt"

func generateConditionalAccountNumber(accountNumberChannel chan int) {
	var accountNumber int 
	accountNumber = 300001

	for  {
		// close the channel after 5 numbers are requested
		if accountNumber > 30001 {
			close(accountNumberChannel)
			return
		}
		accountNumberChannel <- accountNumber
		accountNumber +=1
	}
}

func main() {
	accountNumberChannel := make(chan int)

	go generateConditionalAccountNumber(accountNumberChannel)

	// slice containing account numbers for each customer
	newCustomers := []string{"SMITH", "SINGH", "JONES", "LOPEZ","CLARK", "ALLEN"}

	// get a new account number for each customer
	for _, newCustomer := range newCustomers {
		account, ok := <- accountNumberChannel 
		if !ok {
			fmt.Printf("%s: No number available\n", newCustomer)
		} else {
			fmt.Printf("%s: %d\n", newCustomer, account)
		}
	}
}