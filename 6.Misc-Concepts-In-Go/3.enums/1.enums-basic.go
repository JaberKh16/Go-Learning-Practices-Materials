/*
	Enums Concepts In Go
	====================

	# What Are Enums?

	Go does not have built-in enums like other languages.
	Instead, enums are created using:
	- custom types
	- constants (const)
	- the iota keyword

	This approach provides type safety and readable code.

	# Basic Enum Syntax

	You define a custom type and use const with iota:

		type Status int

		const (
			Pending Status = iota
			Shipped
			Delivered
			Cancelled
		)

	Here:
	- Status is a custom type
	- iota auto-increments values (0, 1, 2, 3...)

	# What is iota?

	iota is a special keyword in Go that:
	- starts at 0
	- increments by 1 for each constant in the block

	Example values:
		Pending   = 0
		Shipped   = 1
		Delivered = 2
		Cancelled = 3

	# Adding String Representation

	Enums are numeric by default, so you often implement String():

		func (s Status) String() string {
			switch s {
			case Pending:
				return "Pending"
			case Shipped:
				return "Shipped"
			case Delivered:
				return "Delivered"
			case Cancelled:
				return "Cancelled"
			default:
				return "Unknown"
			}
		}

	# String-Based Enums (Alternative)

	You can also use string enums:

		type Role string

		const (
			Admin Role = "ADMIN"
			User  Role = "USER"
			Guest Role = "GUEST"
		)

	# Why Use Enums?

	- Improves code readability
	- Prevents invalid values
	- Groups related constants
	- Makes logic easier to manage

	# Key Idea

	Go enums = custom type + const + iota
*/

package main

import (
	"fmt"
)

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


// MarshalJSON for JSON encoding
func (s OrderStatus) MarshalJSON() ([]byte, error)  {
	return []byte(`"` + s.String() + `"`), nil
}

// UnmarshalJSON for JSON decoding (fixed method signature)
func (s *OrderStatus) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	str := string(data)
	if len(str) > 1 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}
	
	// Map string back to enum value
	switch str {
	case "Pending":
		*s = Pending
	case "Received":
		*s = Received
	case "Confirmed":
		*s = Confirmed
	case "Shipped":
		*s = Shipped
	default:
		return fmt.Errorf("invalid OrderStatus: %s", str)
	}
	return nil
}


func main(){
	var status OrderStatus = Shipped
	fmt.Println(status) // value
	fmt.Println(int(status)) // iota value


	var role Role = Admin;
	fmt.Println(role)
}


