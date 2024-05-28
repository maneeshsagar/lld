package service

import (
	"fmt"

	"github.com/maneeshsagar/lld/domain"
)

type BookMyShow struct {
	MovieController   *domain.MovieController
	TheaterController *domain.TheaterController
}

func NewBookMyShow() *BookMyShow {
	return &BookMyShow{
		MovieController:   domain.NewMovieController(),
		TheaterController: domain.NewTheaterController(),
	}
}

func (b *BookMyShow) BookTicket() {

}

func (b *BookMyShow) CancelTicket() {

}

func Initialize() {

}

func (b *BookMyShow) CreateMovies() {

	movie1 := domain.NewMovie(1, "BAHUBALI", 180)
	movie2 := domain.NewMovie(2, "Lapta Ladies", 180)
	movie3 := domain.NewMovie(2, "Madgaon Express", 180)

	b.MovieController.AddMovie(movie1, "Bengaluru")
	b.MovieController.AddMovie(movie2, "Bengaluru")
	b.MovieController.AddMovie(movie3, "Bengaluru")
	b.MovieController.AddMovie(movie3, "Delhi")

}

func (b *BookMyShow) CreateTheater() {

	balhubali, err := b.MovieController.GetMovieByName("BAHUBALI")
	if err != nil {
		fmt.Println("movie not found")
		return
	}

	laptaLadies, err := b.MovieController.GetMovieByName("LAPTA LADIES")
	if err != nil {
		fmt.Println("movie not found")
		return
	}

	theater := &domain.Theater{
		Id:   1,
		City: "Bengaluru",
	}
}

func createShows() (shows []*domain.Show) {
	return
}

func createSceen() {

}
