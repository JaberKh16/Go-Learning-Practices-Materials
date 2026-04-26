// use of errors package
package main

import (
	"errors"
	"fmt"
)

func performDivision(numerator int, denominator int)(int, int, error){
	var err error
	if denominator == 0 {
		// err = fmt.Errorf("division by zero")
		err = errors.New("division by zero")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}

func main() {
	result, remainder, err := performDivision(10, 3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result, remainder)
	}
	
	result, remainder, err = performDivision(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result, remainder)
	}
}
