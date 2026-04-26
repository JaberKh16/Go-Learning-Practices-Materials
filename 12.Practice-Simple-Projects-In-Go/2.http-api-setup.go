package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Todo Model
type Todos struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Generic Response Wrapper
type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// API Setup
type APISetup struct {
	URL string
}

// Perform API Call
func (api *APISetup) performHitTheUrl() (*http.Response, error) {
	resp, err := http.Get(api.URL)
	if err != nil {
		return nil, fmt.Errorf("error hitting url: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp, nil
}

// Option handler
func choseOption(option string, resp *http.Response) ([]Todos, error) {
	defer resp.Body.Close()

	var todos []Todos

	switch option {

	case "1":
		// Read all + Unmarshal
		var body []byte
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading body: %v", err)
		}

		err = json.Unmarshal(body, &todos)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error: %v", err)
		}

	case "2":
		// Decode directly (streaming)
		err := json.NewDecoder(resp.Body).Decode(&todos)
		if err != nil {
			return nil, fmt.Errorf("decode error: %v", err)
		}

	default:
		return nil, fmt.Errorf("invalid option")
	}

	return todos, nil
}

func perfomPostRequest(){
	todo := {
		UserID: 23,
		Title: "Soemthing",
		Completed: true
	}

	// convert the todos struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error marshalling :", err)
	}

	// convert json data sting 
	jsonString := string(jsonData)
	// convert string data to reader
	jsonDataReader := strings.NewReader(jsonString)


	// setup a POST request
	siteUrl := "https://jsonplaceholder,typicode.com/todos"
	response, err := http.Post(siteUrl, "application/json", jsonDataReader)
	if err != nil {
		fmt.Println("error sending request", err)
		return
	}

	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("Request Response: %s, Status: %s",string(data), response.Status)
}


func perfomUpdateRequest(){
	todo := {
		UserID: 2323,
		Title: "Soemthing",
		Completed: true
	}

	// convert the todos struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error marshalling :", err)
	}

	// convert json data sting 
	jsonString := string(jsonData)
	// convert string data to reader
	jsonDataReader := strings.NewReader(jsonString)


	// setup a PUT request
	siteUrl := "https://jsonplaceholder,typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodPut, siteUrl, jsonDataReader)
	if err != nil {
		fmt.Println("error sending request", err)
		return
	}

	// setup the content type
	req.Header.Set("Content-Type", "application/json")

	// send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request: ", err)
	}

	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("Request Response: %s, Status: %s",string(data), res.Status)
}

func performDeleteRequest(){
	// url
	const siteUrl := "https://jsonplaceholder.typicode.com/todos/1"

	// setup a request 
	request, err := http.NewRequest(http.MethodDelete, sitesiteUrl, nil)
	if err != nil {
		fmt.Println("error performing request:DELETE", err)
		return
	}
	
	// send the request
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error sending request: ", err)
		return
	}
	defer response.Body.Close()

	data, _ := ioutil.ReadFile(response.Body)
	fmt.Printf("Request Response: %s, Status: %s",string(data), res.Status)

}

func main() {
	api := APISetup{
		URL: "https://jsonplaceholder.typicode.com/todos",
	}

	resp, err := api.performHitTheUrl()
	if err != nil {
		fmt.Println(err)
		return
	}

	var option string
	fmt.Println("Enter option: 1 = Marshal/Unmarshal, 2 = Decode")
	fmt.Scanln(&option)

	todos, err := choseOption(option, resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print sample
	for i, t := range todos {
		if i >= 5 {
			break
		}
		fmt.Printf("ID: %d | Title: %s | Done: %v\n", t.ID, t.Title, t.Completed)
	}
}