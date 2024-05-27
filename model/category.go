package model

type Category struct {
	ID           uint   `gorm:"primary_key;auto_increment" `
	CategoryName string `gorm:"unique;not null"`
}
