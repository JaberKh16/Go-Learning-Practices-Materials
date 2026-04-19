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
	price       float32
	createdAt   time.Time
	customer: {
		name: "Mr. Robot",
		contact: "+2348-432-2354"
	}
}

// Setter method for price
func (p *Product) setPrice(newPrice float32) {
	p.price = newPrice
}

// Optional: String method for nicer printing
func (p Product) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, SKU: %s, Price: %.2f, CreatedAt: %s",
		p.id, p.name, p.sku, p.price, p.createdAt.Format(time.RFC3339))
}

func initializeStruct(id string, name string, sku string, description string, price float32, createdAt time.Time) *Product {
	instance := Product{
		id:          id,
		name:        name,
		sku:         sku,
		description: description,
		price:       price,
		createdAt:   createdAt,
	}
	return &instance
}

func main() {
	product1 := initializeStruct("Product-1221", "Clock", "clock-m23", "this is clock based product", 23.23, time.Now())

	product1.setPrice(10.22)

	fmt.Println(product1)
}