package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Name string `json:"name"`
	ISBN string `json:"isbn"`
	Year int `json:"year"`
	Director Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
	TotalReleasedMovies int `json:"total_released_movies"`
}

func printData(data interface{}) {
	fmt.Println(data)
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

func appendData(mv *Movie) [] Movie{
	movies = append(movies, *mv)
	return movies
}

func createData(){
	// create a slice of movies
	movie = append(movies, Movie{
		ID: "1",
		Name: "The Godfather",
		ISBN: "1234567890",
		Year: 1972,
		Director: Director{
			Name: "Francis Ford Coppola",
			TotalReleasedMovies: 47,
		},
	})
	appendData(&movie)
}

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder().Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // mux.Vars() contains the request body

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1: ]...)
			break
		}

	}
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
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
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
		}
	}

}

var movies [] Movie

func main() {
	movie := Movie{
		ID: "1",
		Name: "The Godfather",
		ISBN: "1234567890",
		Year: 1972,
		Director: Director{
			Name: "Francis Ford Coppola",
			TotalReleasedMovies: 47,
		},
	}

		
	// print as struct
	printData(movie)
	 
	// print as JSON
	convertStructToJSON(movie)

	// create data
	createData()

	// create a route
	r := mux.NewRouter()


	// handle function
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")


	// run ther server
	const PORT string := 8000
	fmt.Printf("Server run at port: %s", PORT)

	log.Fatal(http.ListenAndServe(":8000", r))


}