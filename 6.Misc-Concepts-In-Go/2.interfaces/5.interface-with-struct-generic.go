package main

import "fmt"

// generic interface
type Printer[T any] interface {
	Print(T)
}

// struct implementing generic interface
type ConsolePrinter[T any] struct{}

func (c ConsolePrinter[T]) Print(val T) {
	fmt.Println("Printing:", val)
}

// generic function using interface
func process[T any](p Printer[T], val T) {
	p.Print(val)
}

func main() {
	intPrinter := ConsolePrinter[int]{}
	stringPrinter := ConsolePrinter[string]{}

	process(intPrinter, 100)
	process(stringPrinter, "hello generics")
}