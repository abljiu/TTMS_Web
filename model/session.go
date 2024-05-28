package model

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	MovieID   uint      `gorm:"movie_id"`
	Movie     Movie     `gorm:"ForeignKey:MovieID"`
	AddressID uint      `gorm:"not null"`
	Address   Address   `gorm:"ForeignKey:AddressID"`
	TheaterID uint      `gorm:"not null"`
	Theater   Theater   `gorm:"ForeignKey:TheaterID"`
	ShowTime  time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
}
