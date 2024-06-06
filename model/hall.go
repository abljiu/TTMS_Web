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
	Seat       string `gorm:"type:varchar(1024)"`
	SeatNum    int
}

//type HallInfo struct {
//	ID         int
//	Name       string
//	TheaterID  int
//	SeatRow    int
//	SeatColumn int
//	Seat       [][]int
//}
