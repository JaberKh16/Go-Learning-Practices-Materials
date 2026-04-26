package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)



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