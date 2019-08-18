package queries

import (
	"fmt"

	mydb "github.com/xenithz/movie-list-api-golang/database"
	types "github.com/xenithz/movie-list-api-golang/movie-types"

	_ "github.com/lib/pq"
)

func GetAllMovies(Movies *types.Movies) error {
	db := mydb.ConnectToDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT * FROM Movies ORDER BY ID ASC`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println("adding movie")
		movie := types.Movie{}
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
		Movies.Movies = append(Movies.Movies, movie)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
