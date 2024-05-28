package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    uint            `gorm:"not null"`
	User      User            `gorm:"ForeignKey:UserID"`
	MovieID   uint            `gorm:"not null"`
	Movie     Movie           `gorm:"ForeignKey:MovieID"`
	AddressID uint            `gorm:"not null"`
	Address   Address         `gorm:"ForeignKey:Address"`
	TheaterID json.RawMessage `gorm:"type:json"`
	Num       uint            `gorm:"not null"`
	ShowTime  string          `gorm:"not null"`
	Type      uint            //1 待支付 2 已支付 3 已完成
	Money     float64
}
