package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User    User  `gorm:"ForeignKey:UserID"`
	UserID  uint  `gorm:"not null"`
	Movie   Movie `gorm:"ForeignKey:MovieID"`
	MovieID uint  `gorm:"not null"`
	Boss    User  `gorm:"ForeignKey:BossID"`
	BossID  uint  `gorm:"not null"`
}
