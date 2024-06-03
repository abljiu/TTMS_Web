package serializer

import (
	"TTMS_Web/model"
)

type Comment struct {
	CommentId    uint   `json:"comment_id"`
	Content      string `json:"content"  `
	UserId       uint   `json:"user_id" `
	RlyId        uint   `json:"rly_id"`
	MovieID      uint   `json:"movie_id"`
	Rate         int    `json:"rate"`
	UpvoteNum    int    `json:"upvote_num"`
	IsSelfUpvote bool   `json:"is_self_upvote"`
	IP           string `json:"ip" `
}

func BuildComment(comment *model.Comment) *Comment {
	return &Comment{
		CommentId:    comment.ID,
		Content:      comment.Content,
		UserId:       comment.UserID,
		RlyId:        uint(comment.RlyID.Int64),
		MovieID:      comment.MovieID,
		Rate:         int(comment.Rate),
		UpvoteNum:    int(comment.UpvoteNum),
		IsSelfUpvote: comment.IsSelfUpvote,
		IP:           comment.IP,
	}
}
func BuildComments(items []*model.Comment) (Comments []*Comment) {
	for _, i := range items {
		Comment := BuildComment(i)
		Comments = append(Comments, Comment)
	}
	return Comments
}
