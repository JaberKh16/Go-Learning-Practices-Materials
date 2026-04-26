package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// --------------------------------
// Interface Type Assertion Demo
// --------------------------------
func interfaceTypeAssertionIntro() {
	fmt.Println("Running Interface Type Assertion....")

	var a interface{} = "some" // define a interface having type string
	var b any = "thing"       // define any having type string i.e interface {}

	c := a.(string) // if declared interface {} or any not type of string it will panic
	fmt.Println(c, b)

	// check on panic situation (safe assertion)
	if c, ok := a.(string); ok {
		fmt.Println("Safe assertion:", c)
	}
	fmt.Println("--------------------------------")
}


func onPrimitiveType[T any](val T) {
	fmt.Printf("primitive handler running, type: %T, simulated value: %s", val, val)
}

func onComplexType[T any](val T) {
	fmt.Printf("complex handler running, type: %T, imulated value: %s", val, val)
}

func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the value: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}


func performOperation(choice int, input string) {
	switch choice {

	case 1:
		// simulate primitive types: int, float, rune, string, bool, byte
		// var checkType interface {} 
		// checkType := input.(string)
		onPrimitiveType(input) // string input

	case 2:
		// simulate complex type
		typeOption := map[string]string{
			"value": input,
		}
		onComplexType(typeOption)

	default:
		fmt.Println("invalid option..")
	}
}


func main() {

	// interface {}  type assertion
	interfaceTypeAssertionIntro()

	// take input => anonymous func
	choiceVal, err := func() (int, error) {
		fmt.Println("Enter choice\n 1:Primitive\n 2:Complex\n")

		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			return 0, fmt.Errorf("error in scanning: %v", err)
		}

		return option, nil
	}()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Step 2: get string input
	input, err := getInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(input)

	// Step 3: perform operation
	performOperation(choiceVal, input)
}