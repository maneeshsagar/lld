package domain

type Movie struct {
	Id       int
	Name     string
	Duration int
}

func NewMovie(id int, name string, duration int) *Movie {
	return &Movie{
		Id:       id,
		Name:     name,
		Duration: duration,
	}
}
