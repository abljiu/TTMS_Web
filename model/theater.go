package model

import "gorm.io/gorm"

type Theater struct {
	gorm.Model
	Name    string
	Address string
	HallNum int
	Movies  []Movie `gorm:"many2many:movie_theaters;"`
}
