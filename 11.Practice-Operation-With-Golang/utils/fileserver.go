package main

import (
	"flag"
	"log"
	"net/http"
)

func createFileServer() {
	// flag package is used for command line interface

	port := flag.String("p", "8000", "port")
	dir := flag.String("d", ".", "dir")
	flag.Parse() // parses the os.Args

	// make http file server
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Printf("Serving %s on http port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":" + *port, nil))
}

func main() {
	// call the file server
	createFileServer()
}