package model

import "gorm.io/gorm"

type Theater struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Address string
	HallNum int
}
