package types

type Movie struct {
	ID            int
	MovieID       int
	MovieTitle    string
	MovieGenre    string
	MovieDirector string
}

type Movies struct {
	Movies []Movie
}
