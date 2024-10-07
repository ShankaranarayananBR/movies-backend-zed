package requests

type RegisterNewMovieRequest struct {
	MovieName   string `json:"movie_name" validate:"required"`
	MovieGenre  string `json:"movie_genre" validate:"required"`
	Description string `json:"description" validate:"required"`
	MovieYear   int64  `json:"movie_year" validate:"required"`
}

type UpdateMovieRequest struct {
	MovieName string `json:"movie_name" validate:"required"`
	MovieYear int64  `json:"movie_year" validate:"required"`
}
