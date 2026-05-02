package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"your_project/generated/models"
)

{{ range . }}

// In-memory storage (replace with DB later)
var {{ .Name }}Store = make(map[int]models.{{ .Name }})
var {{ .Name }}ID = 1

func Create{{ .Name }}(w http.ResponseWriter, r *http.Request) {
	var item models.{{ .Name }}
	json.NewDecoder(r.Body).Decode(&item)

	item.ID = {{ .Name }}ID
	{{ .Name }}Store[item.ID] = item
	{{ .Name }}ID++

	json.NewEncoder(w).Encode(item)
}

func GetAll{{ .Name }}(w http.ResponseWriter, r *http.Request) {
	var list []models.{{ .Name }}
	for _, v := range {{ .Name }}Store {
		list = append(list, v)
	}
	json.NewEncoder(w).Encode(list)
}

func Get{{ .Name }}(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	item, ok := {{ .Name }}Store[id]
	if !ok {
		http.Error(w, "{{ .Name }} not found", 404)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func Update{{ .Name }}(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var item models.{{ .Name }}
	json.NewDecoder(r.Body).Decode(&item)

	item.ID = id
	{{ .Name }}Store[id] = item

	json.NewEncoder(w).Encode(item)
}

func Delete{{ .Name }}(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	delete({{ .Name }}Store, id)
	w.Write([]byte("Deleted"))
}

{{ end }}