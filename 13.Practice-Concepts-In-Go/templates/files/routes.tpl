package routes

import (
	"net/http"
	"your_project/generated/handlers"
)

func RegisterRoutes() {
{{ range . }}
	http.HandleFunc("/{{ .Name | lower }}/create", handlers.Create{{ .Name }})
	http.HandleFunc("/{{ .Name | lower }}/all", handlers.GetAll{{ .Name }})
	http.HandleFunc("/{{ .Name | lower }}/get", handlers.Get{{ .Name }})
	http.HandleFunc("/{{ .Name | lower }}/update", handlers.Update{{ .Name }})
	http.HandleFunc("/{{ .Name | lower }}/delete", handlers.Delete{{ .Name }})
{{ end }}
}