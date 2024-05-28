package domain

type MovieController struct {
	CityVsMovies map[string][]*Movie
	AllMovies    []Movie
}
