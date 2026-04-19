package main

import "fmt"

// define a custom type
type OrderStatus int
type Role string

// enum values => OrderStatus
const (
	Pending OrderStatus= iota
	Received 
	Confirmed
	Shipped
)

// enum values => Role
const (
	Admin Role = "ADMIN"
	User  Role = "USER"
	Guest Role = "GUEST"
)


// custom function
func (s OrderStatus) String() string {
	switch s {
	case Pending:
		return  "Pending"
	case Received:
		return  "Received"
	case Confirmed:
		return  "Confirmed"
	case Shipped:
		return "Shipped"
	default:
		return  "invalid"
	}
}

func main(){
	var status OrderStatus = Shipped
	fmt.Println(status) // value
	fmt.Println(int(status)) // iota value


	var role Role = Admin;
	fmt.Println(role)
}


