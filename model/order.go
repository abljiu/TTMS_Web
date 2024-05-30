package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"ForeignKey:UserID"`
	MovieID   uint    `gorm:"not null"`
	Movie     Movie   `gorm:"ForeignKey:MovieID"`
	TheaterID uint    `gorm:"not null"`
	Theater   Theater `gorm:"ForeignKey:TheaterID"`
	SessionID uint    `gorm:"not null"`
	Session   Session `gorm:"ForeignKey:SessionID"`
	Seat      string  `gorm:"not null"`
	Num       int     `gorm:"not null"`
	Type      uint    //0 待支付 1 已支付 2 已完成
	Money     float64
}
