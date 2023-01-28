package entity

import (
	"gorm.io/gorm"
)

type Passengers struct {
	gorm.Model
	Firstname string  `json:"firstname"`
	Email     string  `json:"email"`
	Lastname  string  `json:"lastname"`
	Ticket    *Ticket `gorm:"-"`
}

type Ticket struct {
	TicketNumber int `json:"ticketnumber"`
}
