package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	address := ":8000"
	mux := http.NewServeMux()

	// instantiated the server handlers
	srv := server.New()

	// routes
	// mux.HandleFunc("/api/users/index",  ).Method("GET")
	// mux.HandleFunc("/api/users/create",  ).Method("POST")
	// mux.HandleFunc("/api/users/{id}",  ).Method("GET")
	// mux.HandleFunc("/api/users/{id}",  ).Method("PUT")
	// mux.HandleFunc("/api/users/{id}",  ).Method("PATCH")
	// mux.HandleFunc("/api/users/{id}",  ).Method("DELETE")


	// server config 
	s := &http.Server{
		Addr: address,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Start server: %v", address)
	log.Fatal(s.ListenAndServe())
}