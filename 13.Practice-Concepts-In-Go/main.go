package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/jaber/go-web/apis"
	"github.com/jaber/go-web/handlers"
	"github.com/jaber/go-web/middlewares"
	"github.com/jaber/go-web/utitlities"
)


func startDefaultHttpServer() error {
	err := http.ListenAndServe(":8000", nil) // nil => DefaultServerMux
	fmt.Println("Server started on http://localhost:8000")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return err
	}
	return nil
}

func startCustomHttpServer(port string) error {
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server started on http://localhost:" + port)
	return server.ListenAndServe()
}


func main() {
	// http.HandleFunc("/", utitlities.HandleIndex)
	// http.HandleFunc("/file/serve", utitlities.HandleServeFile)


	// Start server
	if err := startCustomHttpServer("8000"); err != nil {
		log.Fatal("ListenAndServe:", err)
	}


	time.Sleep(100 * time.Second)

	// working with middleware
	// finalHandler := http.HandlerFunc(middlewares.Final)
	// http.Handle("/", middlewares.Middleware1(middlewares.Middleware2(finalHandler)))


	// call the instance
	// apis.InitializeWithStaticData()

	// setup router
	// router := http.NewServeMux()
	// router.HandleFunc("/product/{id:[0-9]+}", utitlities.pageHandler)

	// router.HandleFunc("/products", productHandler).Host("www.example.com").Methods("GET").Schemes("https")

	// get the resposne
	
	handlers.WorkingWithXMLData()
}


