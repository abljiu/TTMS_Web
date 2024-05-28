package model

type Address struct {
	ID          uint   `gorm:"primary_key"`
	AddressName string `gorm:"unique;not null"`
}
