package main
import (
  "fmt"
  "time"
)


type Address struct {
  addr string
}

type City struct {
  name string
  province string
  zipcode string
}

type User struct {
  Name string
  Age int 
  Gender string
  DOB time.Time
  Address
  City
}

func updateUser(user *User){
	if !checkForValidation(user) {
		updateUser(user)
	}
}

func checkForValidation(user *User) bool {
	if user.Name == "" || user.Age == 0 || user.Gender == "" || user.DOB.IsZero() || user.Address.addr == "" || user.City.name == "" || user.City.province == "" || user.City.zipcode == "" {
		return false
	}
	return true
}

func printUser(user User) {
	fmt.Println(user)
}

func NewUser(name string, age int, gender string, dob time.Time, address Address, city City) *User {
	return &User{
		Name: name,
		Age: age,
		Gender: gender,
		DOB: dob,
		Address: address,
		City: city,
	}
}

func runProgram() {
	// Your program logic here
	for {
		// Your program logic here
		fmt.Println("Running program...")
		time.Sleep(1 * time.Second)

		// switch case for user input
		fmt.Println("Press 'q' to quit")
		var input string
		fmt.Scan(&input)
		switch input {
		case "q":
			return
		case "a":
			// Add new user
			var name string
			var age int
			var gender string
			var dob time.Time
			var addr string
			var city_name string
			var province string
			var zipcode string
			fmt.Scan(&name, &age, &gender, &dob, &addr, &city_name, &province, &zipcode)
			addUser := NewUser(name, age, gender, dob, Address{addr: addr}, City{name: city_name, province: province, zipcode: zipcode})
			
		case "v":
			// View all users
			printUser(user)
		}

		// Add break condition if needed
		// break
	}
}

func staticInit() {
	// initialize city
	city := City {
		name: "Marline",
		province: "mariate 77 AL/O",
		zipcode: "872-2321",
	}

	// initialize user
	user := NewUser("Mr X", 30, "male", time.Now(), Address{addr: "default"}, city)
	
	// initialize user
	// user := User{ 
	// 	Name: "Mr X",
	// 	Age: 30,
	// 	Gender: "male",
	// 	DOB: time.Now(),
	// 	Address:  Address{
	// 		addr: "default",
	// 	},
	// 	City: city,
	// }  
}

func main() {
	staticInit()
	
	// printUser(*user)
  
}