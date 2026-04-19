package main

import "fmt"

func main(){
	
	// struct type variable
	type User struct {
		Name string
		Age  int
	}

	users := make(map[string]User)
	users["u1"] = User{Name: "Alice", Age: 30}

	// print
	fmt.Println(users)
}
