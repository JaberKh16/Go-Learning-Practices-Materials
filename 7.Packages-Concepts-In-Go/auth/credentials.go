package auth

// convention to use Capital
func LogginWithCredentials(username string, password string) (string, string) {
	return username, password
}