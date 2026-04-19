package main

import (
	"fmt"
	"time"
)

type Product struct {
	id string
	name string
	sku string
	description string
	price float32
	createdAt time.Time
}

// dereferencing not required while getting value
func (p Product) getProductInfo() string{
	return fmt.Sprintf("ID: %s, Name: %s, SKU: %s, Description: %s, Price: %.2f, Created: %v", p.id, p.name, p.sku, p.description, p.price, p.createdAt)
}

// change of price requires pointer and struct type auto resolves the dereferencing
func (p *Product) setPrice(price float32){
	p.price = price
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

func main(){

	// custom constructor based function 
	product1 := initializeStruct("Product-1221", "Clock", "clock-m23", "this is clock based product", 23.23, time.Now())
	product1.setPrice(10.22)
	fmt.Println(product1)
}