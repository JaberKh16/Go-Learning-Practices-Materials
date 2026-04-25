package controllers

import "net/http"


func New () http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		message := "welcome to the api"
		w.Write([]byte(message)) // w.Write([]byte())
	}
}