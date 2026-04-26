package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https:s//jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Errorf("some error happend: %v", err.Error())
		return
	}
	// fmt.Printf("Type of response: %T\n", response) // type of response
	// fmt.Println("Response: ", response) // respone itself

	defer response.Body.Close() // close the response body

	// check if the status is pok
	if response.StatusCode != http.StatusOK {
		fmt.Printf("status hits error: ", response.Status)
	}

	// read the content of the response => ioutil.ReadAll() returns array of bytes
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("Error in reading response", err.Error())
		return
	}

	// string([]byte varName) returns the string version of it
	fmt.Println("Content: ", string(content))
}