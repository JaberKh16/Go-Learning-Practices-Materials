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
type ApplicationStatus int 
type JobStatus int 
type HTTPMethod int 

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


// database state
const (
	Applied ApplicationStatus = iota
	Interviewing
	Approved
	Rejected
)

// http method
const (
	GET HTTPMethod = iota 
	POST 
	PUT 
	DELETE
)


// job status queue
const (
	Queued JobStatus = iota
	Running
	Success 
	Failed 
	Retrying
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

func (s OrderStatus) MarshalJSON() ([]byte, error)  {
	return []byte(`"` + s.String() + `"`), nil
}

func (s OrderStatus) Unmarshal() ([]byte, error)  {
	return []byte()
}

func main(){
	var status OrderStatus = Shipped
	fmt.Println(status) // value
	fmt.Println(int(status)) // iota value


	var role Role = Admin;
	fmt.Println(role)
}


