package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	chi "github.com/go-chi/chi"
	types "github.com/xenithz/movie-list-api-golang/movie-types"
	queries "github.com/xenithz/movie-list-api-golang/queries"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := types.Movies{}

	err := queries.GetAllMovies(&Movies)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.MarshalIndent(movies, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func getMovieByID(w http.ResponseWriter, r *http.Request) {
	movies := types.Movies{}
	idString := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic("conversion to int failed in getmoviesbyid")
	}

	err = queries.GetMovieByID(&movies, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.MarshalIndent(movies, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func getMovieByTitle(w http.ResponseWriter, r *http.Request) {
	movies := types.Movies{}
	title := r.URL.Query().Get("title")

	err := queries.GetMovieByTitle(&movies, title)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.MarshalIndent(movies, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func getMovieByGenre(w http.ResponseWriter, r *http.Request) {
	movies := types.Movies{}
	genre := r.URL.Query().Get("genre")

	err := queries.GetMovieByGenre(&movies, genre)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.MarshalIndent(movies, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func getMovieByDirector(w http.ResponseWriter, r *http.Request) {
	movies := types.Movies{}
	director := r.URL.Query().Get("director")

	err := queries.GetMovieByDirector(&movies, director)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.MarshalIndent(movies, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func handleRoutes() {
	router := chi.NewRouter()
	router.Get("/", homePage)
	router.Get("/movies", getAllMovies)
	router.Get("/movies/{id}", getMovieByID)
	router.Get("/movies?title=", getMovieByTitle)
	router.Get("/movies?genre=", getMovieByGenre)
	router.Get("/movies?director=", getMovieByDirector)
	log.Fatal(http.ListenAndServe(":1337", router))
}

func main() {
	handleRoutes()
}
