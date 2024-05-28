package model

import "gorm.io/gorm"

type Hall struct {
	gorm.Model
	SeatNum uint
}
