package main

import "fmt"

// Interface
type PaymentProcessor interface {
	Pay(amount float64)
}

// Implementation 1
type CreditCard struct{}

func (c CreditCard) Pay(amount float64) {
	fmt.Println("Paid using Credit Card:", amount)
}

// Implementation 2
type Bkash struct{}

func (b Bkash) Pay(amount float64) {
	fmt.Println("Paid using Bkash:", amount)
}

// Function using interface
func makePayment(p PaymentProcessor) {
	p.Pay(100)
}

func main() {
	cc := CreditCard{}
	bk := Bkash{}

	makePayment(cc)
	makePayment(bk)
}