package domain

import (
	"slices"

	"github.com/maneeshsagar/lld/errors"
)

type MovieController struct {
	CityVsMovies map[string][]*Movie
	AllMovies    []*Movie
}

func NewMovieController() *MovieController {
	return &MovieController{
		CityVsMovies: make(map[string][]*Movie),
	}
}

func (m *MovieController) AddMovie(movie *Movie, city string) {
	m.AllMovies = append(m.AllMovies, movie)
	m.CityVsMovies[city] = append(m.CityVsMovies[city], movie)
}

func (m *MovieController) RemoveMovie(movie *Movie, city string) {

	if movie == nil {
		return
	}

	movieIndexOfAllMovies := -1
	for i := 0; i < len(m.AllMovies); i++ {
		if m.AllMovies[i].Name == movie.Name {
			movieIndexOfAllMovies = i
		}
	}

	if movieIndexOfAllMovies != -1 {
		slices.Delete(m.AllMovies, movieIndexOfAllMovies, movieIndexOfAllMovies+1)
	}

	moviesIndexFromMap := -1

	for i := 0; i < len(m.CityVsMovies[movie.Name]); i++ {
		if m.CityVsMovies[movie.Name][i].Name == movie.Name {
			moviesIndexFromMap = i
		}
	}

	if movieIndexOfAllMovies != -1 {
		slices.Delete(m.CityVsMovies[movie.Name], moviesIndexFromMap, moviesIndexFromMap+1)
	}

}

func (m *MovieController) GetMovieByName(name string) (*Movie, error) {
	for _, value := range m.AllMovies {
		if value.Name == name {
			return value, nil
		}
	}
	return nil, errors.ErrNotFound
}
