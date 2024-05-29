package model

import "gorm.io/gorm"

type Hall struct {
	gorm.Model
	ID         int
	Name       string
	TheaterID  int
	SeatRow    int
	SeatColumn int
	Seat       string
	SeatNum    int
}
