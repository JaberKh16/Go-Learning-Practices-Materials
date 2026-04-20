package auth

import "fmt"

// convention to use Capital
// func LogginWithCredentials(username string, password string) (string, string) {
// 	return username, password
// }


type Credentials struct {
	Username string
	Password string
}

func LogginWithCredentials(username, password string) (Credentials, error) {
	if username == "" || password == "" {
		return Credentials{}, fmt.Errorf("invalid credentials")
	}

	return Credentials{
		Username: username,
		Password: password,
	}, nil
}