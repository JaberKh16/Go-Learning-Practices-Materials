package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

// generics
func onPrimitiveType[T any](val T) {
	fmt.Printf("primitive handler running, type: %T, simulated value: %s", val, val)
}

func onComplexType[T any](val T) {
	fmt.Printf("complex handler running, type: %T, simulated value: %s", val, val)
}



func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the (type,value): ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func checkSimulatorOption(typeType string, typeVal string) (any, error) {
	switch typeType {
	case "int":
		v, err := strconv.Atoi(typeVal)
		if err != nil {
			return nil, err
		}
		return v, nil

	case "float":
		v, err := strconv.ParseFloat(typeVal, 64)
		if err != nil {
			return nil, err
		}
		return v, nil

	case "byte":
		if len(typeVal) == 0 {
			return nil, fmt.Errorf("empty value for byte")
		}
		return typeVal[0], nil

	case "string":
		return typeVal, nil

	case "bool":
		v, err := strconv.ParseBool(typeVal)
		if err != nil {
			return nil, err
		}
		return v, nil
	case "array":
		arr := strings.Split(typeVal, ",")
		return arr, nil
	case "maps":
		// input format: key1:val1,key2:val2
		m := make(map[string]string)
		pairs := strings.Split(typeVal, ",")
		for _, p := range pairs {
			kv := strings.SplitN(p, ":", 2)
			if len(kv) != 2 {
				return nil, fmt.Errorf("invalid map format")
			}
			m[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
		return m, nil

	case "struct":
		// simulate struct as map
		s := map[string]string{
			"value": typeVal,
		}
		return s, nil

	case "interfaces":
		var i any = typeVal
		return i, nil

	default:
		return nil, fmt.Errorf("unsupported type")
	}
}


func performOperation(choice int, typeType, typeVal string) {
	switch choice {

	case 1:
		// simulate primitive types: int, float, rune, string, bool, byte
		// var checkType interface {} 
		// checkType := input.(string)

		simulatedVal, err := checkSimulatorOption(typeType, typeVal)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		onPrimitiveType(simulatedVal)

	case 2:
		// simulate complex types: array, structs, map, interface
		simulatedVal, err := checkSimulatorOption(typeType, typeVal)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		onComplexType(simulatedVal)

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
	typeType, err := getInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	typeVal, err := getInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(typeType, typeVal)

	// Step 3: perform operation
	performOperation(choiceVal, typeType, typeVal)
}