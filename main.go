package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Movie Struct (Model)
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director Struct
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init Movies var as a slice Movie Struct
var movies []Movie

// Get All Movies
func getMovies(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Get Single Movie
func getMovie(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get params
	params := mux.Vars(router)
	// Loop through books and find with id
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

// Create New Book
func createMovie(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(router.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// Update Movie
func updateMovie(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(router.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// Delete Movie
func deleteMovie(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock Data - @todo - implement DB
	movies = append(movies, Movie{ID: "1", Isbn: "234098", Title: "Movie One", Director: &Director{Firstname: "Elon", Lastname: "Musk"}})
	movies = append(movies, Movie{ID: "2", Isbn: "987456", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Jobs"}})

	// Route Handlers / Endpoints
	router.HandleFunc("/api/movies", getMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/api/movies", createMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
