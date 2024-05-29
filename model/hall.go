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

//type HallInfo struct {
//	ID         int
//	Name       string
//	TheaterID  int
//	SeatRow    int
//	SeatColumn int
//	Seat       [][]int
//}
