package model

import (
	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	ChineseName  string
	EnglishName  string
	CategoryId   json.RawMessage `gorm:"type:json"`
	Area         string
	Duration     string
	ShowTime     string
	Introduction string
	ImgPath      string
	OnSale       bool `gorm:"default:false"`
	Score        float64
	Directors    []Director `gorm:"many2many:movie_directors;"`
	Actors       []Actor    `gorm:"many2many:movie_actors;"`
}
