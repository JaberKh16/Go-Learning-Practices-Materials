package main

import "fmt"

// Interface for numeric operations
type Summable interface {
	Sum(values ...interface{}) int
}

// IntSummer implements Summable for integers
type IntSummer struct{}

func (is *IntSummer) Sum(values ...interface{}) int {
	total := 0
	for _, v := range values {
		if num, ok := v.(int); ok {
			total += num
		}
	}
	return total
}

// FloatSummer for float operations
type FloatSummer struct{}

func (fs *FloatSummer) Sum(values ...interface{}) float64 {
	total := 0.0
	for _, v := range values {
		if num, ok := v.(float64); ok {
			total += num
		}
	}
	return total
}

func main() {
	// Use case: Integer summation
	intSummer := &IntSummer{}
	result := intSummer.Sum(1, 313, 1313, 46, 24, 242)
	fmt.Println(result)

	// Use case: Float summation
	floatSummer := &FloatSummer{}
	result2 := floatSummer.Sum(1.5, 2.5, 3.0)
	fmt.Println(result2)
}
