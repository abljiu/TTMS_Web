package model

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	MovieID uint  `gorm:"movie_id"`
	Movie   Movie `gorm:"ForeignKey:MovieID"`
}
