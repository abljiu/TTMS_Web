package model

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	Name     string
	ImageURL string
	Movies   []Movie `gorm:"many2many:movie_actors;"`
}
