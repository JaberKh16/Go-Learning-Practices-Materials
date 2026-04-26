package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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