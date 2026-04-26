package dto

type UserRegisterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}


type UserRegisterResponse struct {
	ID int `json:"id"`
	Message string `json:"message"`
}


type UserLoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}