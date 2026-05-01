package utitlities

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func HandleIndex(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Current time: ", time.Now())
	fmt.Fprintln(w, "Url Path: ", html.EscapeString(r.URL.Path))
}

func HandleServeFile(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "index.html") // to serve .html file 
}	

func HandleServerDirectory(w http.ResponseWriter, r *http.Request){
	requiredDirectory := "/var/www"
	files := http.FileServer(http.Dir(requiredDirectory))
	http.Handle("/site/", http.StripPrefix("/site", files))
	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		log.Fatal("Listen and server: ", err)
	}
}


func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]
	log.Printf("Product ID: %v\n", productID)

	fileName := productID + ".html"
	log.Printf("File name: %v\n", fileName)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		fileName = "404.html"
		return
	}
	http.ServeFile(w, r, fileName)
}
