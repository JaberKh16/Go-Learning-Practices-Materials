package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	// "github.com/jaber/go-web/apis"
	"github.com/jaber/go-web/handlers"
	// "github.com/jaber/go-web/middlewares"
	// "github.com/jaber/go-web/utitlities"
	"github.com/jaber/go-web/data"
)


func startDefaultHttpServer() error {
	fmt.Println("Server started on http://localhost:8000")
	err := http.ListenAndServe(":8000", nil) // nil => DefaultServerMux
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

	// get the resposne
	handlers.WorkingWithXMLData()


	// work on map
	data.WorkOnMap()


	// handlers setup here
	// http.HandleFunc("/", utitlities.HandleIndex)
	// http.HandleFunc("/file/serve", utitlities.HandleServeFile)
	// http.HandleFunc("/xml/data", handlers.PerformRequest)	


	// Start server
	// if err := startCustomHttpServer("8000"); err != nil {
	// 	log.Fatal("ListenAndServe:", err)
	// }


	// working with middleware
	// finalHandler := http.HandlerFunc(middlewares.Final)
	// http.Handle("/", middlewares.Middleware1(middlewares.Middleware2(finalHandler)))


	// call the instance
	// apis.InitializeWithStaticData()

	// setup router
	// router := http.NewServeMux()
	// router.HandleFunc("/product/{id:[0-9]+}", utitlities.pageHandler)

	// subrouter instance
	// router.HandleFunc("/products", productHandler).Host("www.example.com").Methods("GET").Schemes("https")

	
}


