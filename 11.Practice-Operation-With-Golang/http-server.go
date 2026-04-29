package main

import (
	"log"
	"net/http"
)

// define dto
type Coaster struct {
	Name string `json:"name"`
	Manufacture string `json:"manufacture"`
	ID string `json:"id"`
	InPark string `json:"inPark"`
	Height int  `json:"height"`
}

// define type
type coasterHandlers struct {
	store map[string]Coaster
}



func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request){

}

// constructor
func newCoasterHandlers() *coasterHandlers {
	return  &coasterHandlers{
		store: map[string]Coaster{
			staticData()
		},
	}
}

func staticData() (map[string]Coaster){
	data := Coaster{
		"id1": Coaster{
			Name: "Fury 325",
			Height: 99,
			ID: "id1",
			InPark: "Carowinds",
			Manufacture: "B+M",
		}
	}
}

func main() {

	// define the instance
	coasterHandlers := newCoasterHandlers()

	// handlers
	http.HandleFunc("/coasters", coasterHandlers.get)


	const PORT = "8000"
	err := http.ListenAndServe(":" + PORT, nil)
	if err != nil {
		panic(err)
		log.Printf("failed to start the server: %v", err)
	}


}