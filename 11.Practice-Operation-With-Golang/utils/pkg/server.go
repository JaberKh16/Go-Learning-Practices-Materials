package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Request/Response struct
type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Internal storage
type userinfo struct {
	email string
	age   int
}

// Server
type Server struct {
	users map[string]userinfo
	mu    sync.RWMutex // concurrency safe
}

// Constructor
func New() *Server {
	return &Server{
		users: make(map[string]userinfo),
	}
}

// -------------------- HANDLERS --------------------

// Index
func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User API Running 🚀"))
}

// CREATE
func (s *Server) HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var u user
	if err := json.Unmarshal(body, &u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	s.users[u.Name] = userinfo{
		email: u.Email,
		age:   u.Age,
	}
	s.mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user created"})
}

// READ (single)
func (s *Server) HandleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")

	s.mu.RLock()
	u, ok := s.users[name]
	s.mu.RUnlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := user{
		Name:  name,
		Email: u.email,
		Age:   u.age,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// READ ALL
func (s *Server) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	var list []user
	for name, u := range s.users {
		list = append(list, user{
			Name:  name,
			Email: u.email,
			Age:   u.age,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

// UPDATE
func (s *Server) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var u user
	if err := json.Unmarshal(body, &u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.users[name]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s.users[name] = userinfo{
		email: u.Email,
		age:   u.Age,
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "user updated"})
}

// DELETE
func (s *Server) HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")

	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.users[name]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(s.users, name)

	json.NewEncoder(w).Encode(map[string]string{"message": "user deleted"})
}

// -------------------- MAIN --------------------

func main() {
	srv := New()

	http.HandleFunc("/", srv.HandleIndex)
	http.HandleFunc("/create", srv.HandleCreate)
	http.HandleFunc("/user", srv.HandleGet)
	http.HandleFunc("/users", srv.HandleGetAll)
	http.HandleFunc("/update", srv.HandleUpdate)
	http.HandleFunc("/delete", srv.HandleDelete)

	port := 8080
	log.Printf("Server running at http://localhost:%d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}