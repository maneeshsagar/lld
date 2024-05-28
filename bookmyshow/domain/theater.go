package domain

type Theater struct {
	Id      int
	City    string
	Screens []*Screen
	Shows   []*Show
}
