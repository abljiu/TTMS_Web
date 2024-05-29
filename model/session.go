package model

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	MovieID   uint      `gorm:"movie_id"`
	Movie     Movie     `gorm:"ForeignKey:MovieID"`
	TheaterID uint      `gorm:"not null"`
	Theater   Theater   `gorm:"ForeignKey:id"`
	HallID    uint      `gorm:"not null"`
	Hall      Hall      `gorm:"ForeignKey:id"`
	ShowTime  time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
}
