package main

import (
	"fmt"
	"time"
)

type Customer struct {
	name    string
	contact string
}

type Product struct {
	id          string
	name        string
	sku         string
	description string
	price       float64
	createdAt   time.Time
	customer    Customer
}

// Setter method for price
func (p *Product) setPrice(newPrice float64) {
	if newPrice < 0 {
		fmt.Println("price cannot be negative")
		return
	}
	p.price = newPrice
}

// String method for nicer printing
func (p Product) String() string {
	return fmt.Sprintf(
		"ID: %s, Name: %s, SKU: %s, Price: %.2f, Customer: %s, CreatedAt: %s",
		p.id,
		p.name,
		p.sku,
		p.price,
		p.customer.name,
		p.createdAt.Format(time.RFC3339),
	)
}

// Constructor-like function
func newProduct(id string, name string, sku string, description string, price float64, customer Customer) *Product {
	return &Product{
		id:          id,
		name:        name,
		sku:         sku,
		description: description,
		price:       price,
		createdAt:   time.Now(),
		customer:    customer,
	}
}

func main() {

	// intialized the customer struct
	customer := Customer{
		name:    "John Doe",
		contact: "123456789",
	}

	product1 := newProduct(
		"Product-1221",
		"Clock",
		"clock-m23",
		"this is clock based product",
		23.23,
		customer,
	)

	product1.setPrice(10.22)

	fmt.Println(product1)
}