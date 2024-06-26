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

// ExitOrNorByUserID  根据UserID 查询用户名是否存在
func (dao *UserDao) ExitOrNorByUserID(userID string) (user *model.User, exit bool, err error) {
	var users []model.User
	err = dao.DB.Model(&model.User{}).Where("id=?", userID).Find(&users).Error
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

func (dao *UserDao) GetUserByID(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

func (dao *UserDao) UpdateUserByID(uid uint, user *model.User) error {
	err := dao.DB.Model(&model.User{}).Where("id=?", uid).Updates(&user).Error
	return err
}

func (dao *UserDao) GetNewID() (uid uint, err error) {
	err = dao.DB.Model(&model.User{}).Order("id DESC").Limit(1).Pluck("id", &uid).Error
	return uid, err
}
