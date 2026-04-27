package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	address := ":8000"
	mux := http.NewServeMux()
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