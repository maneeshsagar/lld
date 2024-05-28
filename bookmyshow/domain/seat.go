package domain

type Seat struct {
	Id       int
	Row      int
	Category string
}

func NewSeat(id int, row int, category string) *Seat {
	return &Seat{
		Id:       id,
		Row:      row,
		Category: category,
	}
}
