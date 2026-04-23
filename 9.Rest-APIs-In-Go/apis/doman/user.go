package domain

type Address struct {
  addr string
}

type City struct {
  name string
  province string
  zipcode string
}


type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Address   Address
	City      City
}
