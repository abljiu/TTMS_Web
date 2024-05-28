package model

type Theater struct {
	ID          uint    `gorm:"primary_key;auto_increment" `
	TheaterName string  `gorm:"unique;not null"`
	SeatNum     int     `gorm:"not null"`
	AddressID   uint    `gorm:"not null"`
	Address     Address `gorm:"ForeignKey:AddressID"`
}
