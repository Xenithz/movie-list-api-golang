package queries

import (
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

func GetMovieByID(Movies *types.Movies, id int) error {
	db := mydb.ConnectToDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT * FROM movies WHERE movie_id = $1`, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
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

func GetMovieByTitle(Movies *types.Movies, title string) error {
	db := mydb.ConnectToDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT * FROM movies WHERE movie_title = $1`, title)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
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

func GetMovieByGenre(Movies *types.Movies, genre string) error {
	db := mydb.ConnectToDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT * FROM movies WHERE movie_genre = $1`, genre)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
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

func GetMovieByDirector(Movies *types.Movies, director string) error {
	db := mydb.ConnectToDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT * FROM movies WHERE movie_director = $1`, director)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
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
