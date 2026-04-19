/*
Generic Concepts In Go
=======================

# What Are Generics?

Generics allow you to write reusable and type-safe code by using
type parameters instead of specific concrete types.

Instead of writing separate functions for int, string, etc.,
you write one function that works with any type.

# Generic Syntax

	Generics use square brackets [] to define type parameters.

	Example:

		func printSlice[T any](items []T)

	Here:
	- T is a type parameter
	- any is a constraint (means any type is allowed)

	# Type Constraints

	Constraints define what types are allowed.

	1. any (most flexible)
		T can be any type.

	2. interface{}
		Same as "any" (older style, still valid).

	3. Type unions
		Restrict T to specific types:

		func example[T int | string | bool](value T)

	Only int, string, or bool are allowed.

	# Generic Structs

	You can also define structs with generics:

		type Stack[T any] struct {
			elements []T
		}

	# Instantiating Generics

	You must specify the type when using them:

		stack := Stack[string]{
			elements: []string{"Go", "Java", "Python"},
		}

	# Why Use Generics?

	- Avoid code duplication
	- Maintain type safety
	- Write cleaner and reusable code

	# Key Idea

	Write once → use with many types safely.
*/

package main

import "fmt"

// define generic type struct
type Stack[T any] struct {
	elements []T
}

// generic function => via any
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// generic function => via interface
func printInterfaceSlice[T interface{}](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// scoped generic function => via specified type => type | type
func printScopedIntefaceSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// generic function => via comparables(any primitive types support)
func printThroughComparablesSlice[T comparable](items []T){
	for _, item := range items {
		fmt.Println(item)
	}
}

func printInfo[T any](s Stack[T]) string {
	result := fmt.Sprintf("Stack elements: %v", s.elements)
	fmt.Println(result)
	return result
}

func main() {
	/*
		// using [T any]
		printSlice([]int{1, 12, 32, 45})
		printSlice([]string{"C++", "JavaScript", "TypeScript"})

		// using [T interface{}]
		printInterfaceSlice([]int{1, 12, 32, 45})
		printInterfaceSlice([]string{"C++", "JavaScript", "TypeScript"})

		// using [T type | type | type]
		printScopedIntefaceSlice([]int{1, 12, 32, 45})
		printScopedIntefaceSlice([]string{"C++", "JavaScript", "TypeScript"})
	*/

	// use a generic type struct
	stack := Stack[string]{
		elements: []string{"C++", "JavaScript", "TypeScript"},
	}
	printInfo(stack)
}