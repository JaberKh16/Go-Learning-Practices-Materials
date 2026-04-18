package main

func main(){
	
	// struct type variable
	type User struct {
		Name string
		Age  int
	}

	users := make(map[string]User)
	users["u1"] = User{Name: "Alice", Age: 30}

}
