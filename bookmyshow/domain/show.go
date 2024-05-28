package domain

import "time"

type Show struct {
	Id          int
	Movie       *Movie
	Screen      *Screen
	BookedSeats []int
	StartTime   time.Time
}

func NewShow(id int, movie *Movie, screen *Screen, seats []int, startTime time.Time) *Show {
	return &Show{
		Id:          id,
		Movie:       movie,
		Screen:      screen,
		BookedSeats: seats,
		StartTime:   startTime,
	}
}
