package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// structs
type Movie struct {
	ID string `json:"id"`
	Name string `json:"name"`
	ISBN string `json:"isbn"`
	Year int `json:"year"`
	Genre string `json:"genre"`
	Director Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
	TotalReleasedMovies int `json:"total_released_movies"`
}

var movies [] Movie


// -------- Utitlities --------------
func initializeWithStaticData() *Movie {
	movie := Movie{
		ID:    strconv.Itoa(rand.Intn(1000)),
		Name:  "The Godfather",
		ISBN:  strconv.Itoa(rand.Intn(1000000)),
		Year:  1972,
		Genre: "Crime",
		Director: Director{
			Name:                "Francis Ford Coppola",
			TotalReleasedMovies: 47,
		},
	}
	return &movie
}

func convertStructToJSON(data interface{}) {
	// jsonData, _ := json.Marshal(data)
	// fmt.Println(string(jsonData))
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}

func appendData(mv *Movie) {
	movies = append(movies, *mv)
}

func printData(data interface{}) {
	fmt.Println(data)
}



// ---------------- APIs Func Handlers ----------------
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}



func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000)) // strconv.Itoa(int): string
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1: ]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	// if wanted direct motify message
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func patchUpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {

			// decode into a map for partial updates
			var updates map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// apply updates safely
			if name, ok := updates["name"].(string); ok {
				movies[index].Name = name
			}

			if isbn, ok := updates["isbn"].(string); ok {
				movies[index].ISBN = isbn
			}

			if year, ok := updates["year"].(float64); ok {
				movies[index].Year = int(year)
			}

			if genre, ok := updates["genre"].(string); ok {
				movies[index].Genre = genre
			}

			if directorData, ok := updates["director"].(map[string]interface{}); ok {
				if name, ok := directorData["name"].(string); ok {
					movies[index].Director.Name = name
				}
				if total, ok := directorData["total_released_movies"].(float64); ok {
					movies[index].Director.TotalReleasedMovies = int(total)
				}
			}

			json.NewEncoder(w).Encode(movies[index])
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // mux.Vars() contains the request body

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1: ]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}



func main() {
	
	// initialize with data
	movie := initializeWithStaticData()
		
	// // print as struct
	// printData(movie)
	 
	// // print as JSON
	// convertStructToJSON(movie)

	// append pass the movie struct reference
	appendData(movie)



	// create a mux route
	r := mux.NewRouter()


	// handle function
	r.HandleFunc("/api/movies", getMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/api/movies", createMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", patchUpdateMovie).Methods("PATCH")
	r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")


	// run ther server
	const PORT  = "8000"
	fmt.Printf("Server run at port: %s", PORT)

	// log.Fatal(http.ListenAndServe(":8000", r))
	log.Fatal(http.ListenAndServe(":" +  PORT, r))


}