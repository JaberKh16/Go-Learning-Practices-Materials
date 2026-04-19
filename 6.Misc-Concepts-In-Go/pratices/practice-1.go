package main

import "fmt"

type PaymentMethod int

const (
	Cash PaymentMethod = iota
	Card
	MobileBanking
)

func (p PaymentMethod) String() string {
	switch p {
	case Cash:
		return "Cash"
	case Card:
		return "Card"
	case MobileBanking:
		return "Mobile Banking"
	default:
		return "Unknown"
	}
}

type Order struct {
	id     string
	method PaymentMethod
}

func main() {
	order := Order{
		id:     "ORD-101",
		method: MobileBanking,
	}

	fmt.Printf("Order %s paid via %s\n", order.id, order.method)
}