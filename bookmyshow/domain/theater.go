package domain

type Theater struct {
	Id      int
	Address string
	Screens []Screen
	Shows   []Show
}
