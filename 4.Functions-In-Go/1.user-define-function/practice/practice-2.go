package main

import "fmt"

func performDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := performDivision(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}



