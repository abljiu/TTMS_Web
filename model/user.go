package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string
	Money          uint
}

const (
	PasswordCost         = 12              //密码加密难度
	Normal        string = "normal"        //普通用户
	Administrator string = "administrator" //管理员
	Conductor     string = "conductor"     //售票员
)

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func (user *User) SetRoot() error {
	user.PasswordDigest = Administrator
	return nil
}
