package model

import (
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	gorm.Model
	ChineseName  string
	EnglishName  string
	CategoryId   json.RawMessage `gorm:"type:json"`
	Area         string
	Duration     time.Duration
	ShowTime     time.Time
	Introduction string
	ImgPath      string
	OnSale       bool `gorm:"default:false"`
	Score        float64
	Sales        int64      `json:"sales"`
	Directors    []Director `gorm:"many2many:movie_directors;"`
	Actors       []Actor    `gorm:"many2many:movie_actors;"`
}
