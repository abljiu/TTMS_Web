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
	AddressID uint    `gorm:"not null"`
	Address   Theater `gorm:"ForeignKey:TheaterID"`
	TheaterID uint    `gorm:"not null"`
	Theater   Hall    `gorm:"ForeignKey:TheaterID"`
	Seat      string  `gorm:"not null"`
	Num       uint    `gorm:"not null"`
	ShowTime  string  `gorm:"not null"`
	Type      uint    //0 待支付 1 已支付 2 已完成
	Money     float64
}
