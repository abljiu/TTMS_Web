package serializer

import (
	"TTMS_Web/conf"
	"TTMS_Web/model"
)

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	NickName string `json:"nick_name"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
	Money    uint   `json:"money"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:       user.ID,
		Email:    user.Email,
		NickName: user.NickName,
		Status:   user.Status,
		Avatar:   conf.Config_.Path.Host + conf.Config_.Service.HttpPort + conf.Config_.Path.AvatarPath + user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
		Money:    user.Money,
	}
}
