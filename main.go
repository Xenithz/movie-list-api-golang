package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	types "github.com/xenithz/movie-list-api-golang/movie-types"
	queries "github.com/xenithz/movie-list-api-golang/queries"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	Movies := types.Movies{}

	err := queries.GetAllMovies(&Movies)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.MarshalIndent(Movies, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func handleRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/movies", getAllMovies)
	log.Fatal(http.ListenAndServe(":1337", nil))
}

func main() {
	handleRoutes()
}
