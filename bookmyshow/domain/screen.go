package domain

type Screen struct {
	Seats []*Seat
}

func NewScreen(seats []*Seat) *Screen {
	return &Screen{
		Seats: seats,
	}
}
