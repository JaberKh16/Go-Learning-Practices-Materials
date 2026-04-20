package main

import (
	"log"
	"net/http"

	config "github.com/jaber/rest-apis/internal"
)

func main(){
	// load config
	cfg := config.MustLoad()

	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to the api"))
	})

	// setup server
	server := http.Server {
		Addr: cfg.Address,
		Handler: router,
	}

	// start listeing
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %s", err.Error())
	}
}