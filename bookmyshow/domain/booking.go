package domain

type Booking struct {
	Show        *Show
	BookedSeats []*Seat
	Payment     *Payment
}
