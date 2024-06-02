package model

import "gorm.io/gorm"

type MovieThreat struct {
	gorm.Model
	MovieId  uint
	ThreatId uint
}
