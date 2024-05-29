package model

import "gorm.io/gorm"

type Hall struct {
	gorm.Model
	SeatNum    int
	Seat       string
	Name       string
	TheaterID  uint
	SeatRow    uint
	SeatColumn uint
}
