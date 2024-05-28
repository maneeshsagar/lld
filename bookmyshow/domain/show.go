package domain

import "time"

type Show struct {
	Id          int
	Movie       Movie
	Screen      Screen
	BookedSeats []int
	StartTime   time.Time
}
