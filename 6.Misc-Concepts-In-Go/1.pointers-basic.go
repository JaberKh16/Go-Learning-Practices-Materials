/*

	Pointers Concepts In Go
	=======================

	# What Are Pointers?

	A pointer is a variable that stores the memory address of another variable.
	Instead of holding a value directly, it "points" to where the value is stored in memory.

	# Pointer Syntax

	Go uses the following operators:

	- &  → gives the memory address of a variable
	- *  → dereferences a pointer (gets the value at the address)

	Example:

		x := 10
		p := &x   // p stores the address of x
		fmt.Println(*p) // prints 10

	# Declaring a Pointer

		type intPointer *int

	Or directly:

		var p *int

	Initially, a pointer has a nil value if not assigned.

	# Dereferencing a Pointer

	Dereferencing means accessing or modifying the value stored at the memory address:

		x := 5
		p := &x

		*p = 10  // modifies original value of x

	Now x becomes 10.

	# Passing by Value vs Passing by Pointer

	## Passing by Value (default)

		func updateValue(x int) {
			x = 100
		}

	Original variable does NOT change.

	## Passing by Pointer

		func updateValue(x *int) {
			*x = 100
		}

	Original variable changes because memory address is used.

	# Why Use Pointers?

	Pointers are used when you want:

	1. To modify original data inside functions
	2. To avoid copying large structs (performance optimization)
	3. To share data between functions efficiently
	4. To represent optional values (nil checks)

	# Pointer with Structs

		type User struct {
			name string
			age  int
		}

		func updateUser(u *User) {
			u.age = 30
		}

	Structs are commonly passed using pointers to avoid copying.

	# Nil Pointers

	A pointer that is not initialized is nil:

		var p *int

	Dereferencing nil causes runtime panic.

	Always check:

		if p != nil {
			fmt.Println(*p)
		}

	# Pointer Receiver in Methods

		func (u *User) updateName(name string) {
			u.name = name
		}

	Pointer receivers allow methods to modify the original struct.

	# Key Idea

	Pointers allow direct memory access, enabling efficient and mutable operations in Go.
*/

package main

import "fmt"

func changeByValue(value int){
	value = 10
	fmt.Println("Changed value: ", value)
}

func changeByReference(value *int){
	*value = 10 // derefence the pointer
	fmt.Println("Changed value: ", value)
}

func main(){
	value :=12
	changeByValue(value) // pass the value
	fmt.Println("Value is: ", value)
	fmt.Println("Address is: ", &value)


	changeByReference(&value) // pass hhe reference
	fmt.Println("Value is: ", value)
	fmt.Println("Address is: ", &value)
}