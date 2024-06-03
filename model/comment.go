package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content      string
	UserID       uint          `gorm:"user_id"`
	User         User          `gorm:"ForeignKey:UserID"`
	RlyID        sql.NullInt64 `gorm:"rly_id"`
	Comment      *Comment      `gorm:"ForeignKey:RlyID"`
	MovieID      uint          `gorm:"movie_id"`
	Movie        Movie         `gorm:"ForeignKey:MovieID"`
	Rate         uint          `gorm:"rate,omitempty"`
	IP           string        `gorm:"not null"`
	UpvoteNum    uint          `gorm:"upvote_num,omitempty"`
	IsSelfUpvote bool          `gorm:"is_self_upvote"`
}
