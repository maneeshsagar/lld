package domain

type Booking struct {
	Show        *Show
	BookedSeats []*Seat
	Payment     *Payment
}

func NewBooking(show *Show, seats []*Seat, payment *Payment) *Booking {
	return &Booking{
		Show:        show,
		BookedSeats: seats,
		Payment:     payment,
	}
}
