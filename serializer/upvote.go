package serializer

import "TTMS_Web/model"

type Upvote struct {
	UpvoteId  uint `json:"id"`
	UserId    uint `json:"user_id"`
	CommentId uint `json:"comment_id"`
}

func BuildUpvote(upvote *model.Upvote) *Upvote {
	return &Upvote{
		UpvoteId:  upvote.ID,
		UserId:    upvote.UserID,
		CommentId: upvote.CommentID,
	}
}
func BuildUpVotes(items []*model.Upvote) (upvotes []*Upvote) {
	for _, i := range items {
		upvote := BuildUpvote(i)
		upvotes = append(upvotes, upvote)
	}
	return upvotes
}
