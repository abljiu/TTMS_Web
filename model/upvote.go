package model

import "gorm.io/gorm"

type Upvote struct {
	gorm.Model
	UserID    uint    `gorm:"user_id"`
	User      User    `gorm:"ForeignKey:UserID"`
	CommentID uint    `gorm:"comment_id"`
	Comment   Comment `gorm:"ForeignKey:CommentID"`
}
