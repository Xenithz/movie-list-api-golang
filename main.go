package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mydb "github.com/xenithz/movie-list-api-golang/database"

	_ "github.com/lib/pq"
)

type movie struct {
	ID            int
	MovieID       int
	MovieTitle    string
	MovieGenre    string
	MovieDirector string
}

type movies struct {
	Movies []movie
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := movies{}

	err := moviesQuery(&movies)
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

func moviesQuery(movies *movies) error {
	db := mydb.ConnectToDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT * FROM movies ORDER BY ID ASC`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println("adding movie")
		movie := movie{}
		err = rows.Scan(
			&movie.ID,
			&movie.MovieID,
			&movie.MovieTitle,
			&movie.MovieGenre,
			&movie.MovieDirector,
		)
		if err != nil {
			return err
		}
		movies.Movies = append(movies.Movies, movie)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func handleRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/movies", getAllMovies)
	log.Fatal(http.ListenAndServe(":1337", nil))
}

func main() {
	handleRoutes()
}
