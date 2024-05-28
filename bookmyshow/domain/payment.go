package domain

type Payment struct {
	Status string
}

func NewPayment() *Payment {
	return &Payment{}
}

func (p *Payment) UpdateStatus(status string) {
	p.Status = status
}
