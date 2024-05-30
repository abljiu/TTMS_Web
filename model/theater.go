package model

import "gorm.io/gorm"

type Theater struct {
	gorm.Model
	Name      string
	Address   string
	Telephone string
}
