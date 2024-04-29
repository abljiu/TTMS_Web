package dao

import (
	"TTMS_Web/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExitOrNorByUserName 根据UserName 查询用户名是否存在
func (dao *UserDao) ExitOrNorByUserName(userName string) (user *model.User, exit bool, err error) {
	var users []model.User
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&users).Error
	if len(users) == 0 {
		return nil, false, err
	}
	if err != nil {
		return &users[0], false, err
	}
	return &users[0], true, nil
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}
