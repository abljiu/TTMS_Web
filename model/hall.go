package model

import (
	"gorm.io/gorm"
)

type Hall struct {
	gorm.Model
	Name       string
	TheaterID  uint    `gorm:"theater_id"`
	Theater    Theater `gorm:"ForeignKey:TheaterID"`
	SeatRow    int
	SeatColumn int
	Seat       string
	SeatNum    int
}
