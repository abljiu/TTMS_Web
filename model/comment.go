package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content      string
	UserID       uint     `gorm:"user_id"`
	User         User     `gorm:"ForeignKey:UserID"`
	RlyID        uint     `gorm:"rly_id"`
	Comment      *Comment `gorm:"ForeignKey:RlyID"`
	MovieID      uint     `gorm:"movie_id"`
	Movie        Movie    `gorm:"ForeignKey:MovieID"`
	Rate         int
	IP           string `gorm:"not null"`
	UpvoteNum    int
	IsSelfUpvote bool
}
