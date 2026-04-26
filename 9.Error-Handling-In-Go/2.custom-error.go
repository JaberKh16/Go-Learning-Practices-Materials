package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error occurred due to %v", e.data)
}

func performOperation(num1, num2 int) (int, error) {
	if num1 < 0 && num2 < 0 {
		// creating a 
		return  0, errors.New("values can't be zero or less than zero")
	}
	return num1 + num2, nil
}

func main(){
	result, err := performOperation(10, -1)
	if err != nil {
		// fmt.Println(err.Error())
		// err := fmt.Errorf("error is %v", err.Error())
		// fmt.Println(err)

		// custom error
		var customErr CustomError
		customErr.message = "can't assign zero value"
		fmt.Println(customErr.Error())
		return
	}
	fmt.Println(result)
}