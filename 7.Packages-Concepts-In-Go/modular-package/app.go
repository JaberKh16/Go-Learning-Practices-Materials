package main

import (
	"fmt"

	"github.com/jaber/modules_package/auth"
)

func main(){
	username, pass := auth.LogginWithCredentials("Jk", "password")
	fmt.Println(username, pass)
}
