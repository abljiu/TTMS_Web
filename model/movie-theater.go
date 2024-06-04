package model

import "gorm.io/gorm"

type MovieTheater struct {
	gorm.Model
	MovieID   uint    `gorm:"not null"`
	Movie     Movie   `gorm:"ForeignKey:MovieID"`
	TheaterID uint    `gorm:"not null"`
	Theater   Theater `gorm:"ForeignKey:TheaterID"`
}
