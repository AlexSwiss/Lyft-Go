package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {

}

//get single movie
func getMovie(w http.ResponseWriter, r *http.Request) {

}

//add new movies
func addMovie(w http.ResponseWriter, r *http.Request) {

}

//delete movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {

}

//edit movie
func editMovie(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//initialize router
	r := mux.NewRouter()

	//router endpoints
	r.HandleFunc("/api/movies", getMovies).Methods("GET")
	r.HandleFunc("/api/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/api/movies", addMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/api/movies/{id}", editMovie).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
