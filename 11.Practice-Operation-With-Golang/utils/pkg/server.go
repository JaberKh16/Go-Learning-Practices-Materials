package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Age string `json:"age"`
}

type userinfo struct {
	email string
	age int
}

// define server struct => HTTP Server
type Server struct {
	users map[string]userinfo
}

// constructor for the server
func New() *Server {
	return &Server{}
}


func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Index"))
}

func (s *Server) HandleCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Create"))

	switch r.Method {
	case http.MethodPost, http.MethodPut:
		// check the content type 
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		// unmarshal 
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("couldn't read request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError) // HTTP 500
			return
		}
		defer r.Body().Close()

		var u user
		err = json.Unmarshal(body, &u)
		if err != nil {
			log.Printf("could not unmarshal the request body: %v", err)
			w.WriteHeader(http.StatusBadRequest) // HTTP 400
			return
		}
		log.Printf("Create User: %v", u.Name)
		s.users[u.Name] = userinfo{
			email: u.Email,
			age: u.Age,
		}

	default: 
		w.WriteHeader(http.StatusMethodNotAllowed) // HTTP 415
	}
}


func (s *Server) HandleReadList(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case http.MethodGet:
		name := r.URL.Query().Get("name")
		u, ok := s.users[name]
		if !ok {
			w.WriteHeader(http.StatusOK)
			return
		} 
		ret := user {
			Name: name,
			Email: u.email,
			Age: u.age,
		}		
		msg, err := json.Marshal(ret)
		if err != nil {
			log.Printf("could not marshal request: %v", err)
			w.WriteHeader(http.StatusInternalServerError) // HTTP 500
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(msg)
	default: 
		w.WriteHeader(http.StatusMethodNotAllowed) // HTTP 415
	}
	
	
}

func (s *Server) HandleGetById(w http.ResponseWriter, r *http.Request) {
	
}