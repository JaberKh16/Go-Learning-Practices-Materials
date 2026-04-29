package utils_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
			// return
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
