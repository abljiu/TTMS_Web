package service

import (
	"TTMS_Web/conf"
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	NickName      string `json:"nick_name" form:"nick_name"`
	UserID        string `json:"user_id" form:"user_id"`
	Password      string `json:"password" form:"password"`
	Email         string `json:"email" form:"email"`
	OperationType uint   `json:"operation_type" form:"operation_type"` //1 绑定邮箱 2 解绑邮箱 3 改密码
	Status        string `json:"status" form:"status"`
}

// Register 注册逻辑
func (service *Service) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success

	if service.Password == "" {
		code = e.ErrorNoPassword
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//绑定邮箱
	var address string
	var notice *model.Notice
	token, err := util.GenerateEmailToken(0, 1, service.NickName, service.Email, service.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeByType(1)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	address = conf.Config_.Email.ValidEmail + token //发送方
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.Config_.Email.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "TTMS_Web")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.Config_.Email.SmtpHost, 465, conf.Config_.Email.SmtpEmail, conf.Config_.Email.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    "请验证邮箱",
		Data:   serializer.BuildUser(&user),
	}
}

// Login  登陆逻辑
func (service *Service) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success

	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExitOrNorByUserID(service.UserID)
	if !exist || err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//http 无状态(认证，让对方带上token)
	token, err := util.GenerateToken(user.ID, user.Status)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}
}

// Update 用户修改信息
func (service *Service) Update(ctx context.Context, uid uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByID(uid)
	//修改用户昵称
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserByID(uid, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// Post  上传头像
func (service *Service) Post(ctx context.Context, uid uint, file multipart.File) serializer.Response {
	code := e.Success
	var user *model.User
	var err error
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByID(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//保存图片到本地
	IdStr := strconv.FormatUint(uint64(user.ID), 10)
	path, err := UploadAvatarToLocalStatic(file, uid, IdStr)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserByID(uid, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// Send 发送邮件
func (service *Service) Send(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var address string
	var notice *model.Notice
	token, err := util.GenerateEmailToken(uid, service.OperationType, service.NickName, service.Email, service.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeByType(service.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	address = conf.Config_.Email.ValidEmail + token //发送方
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.Config_.Email.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "TTMS_Web")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.Config_.Email.SmtpHost, 465, conf.Config_.Email.SmtpEmail, conf.Config_.Email.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   token,
	}
}

// Valid 验证邮箱
func (service *Service) Valid(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	var nickName string
	code := e.Success

	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthTokenTimeout
		} else {
			nickName = claims.Nickname
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user := &model.User{
		NickName: nickName,
		Status:   model.Normal,
		Avatar:   "avatar.JPG",
		Money:    0,
	}

	if operationType == 1 {
		user.Email = email
		//如果是首次绑定 说明是注册
		if userId == 0 {
			//密码加密
			userDao := dao.NewUserDao(ctx)
			if err := user.SetPassword(service.Password); err != nil {
				code = e.ErrorFailEncryption
				return serializer.Response{
					Status: code,
					Msg:    e.GetMsg(code),
				}
			}
			//创建用户
			err := userDao.CreateUser(user)
			if err != nil {
				code = e.Error
			}
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   serializer.BuildUser(user),
			}
		}
	} else if operationType == 2 {
		user.Email = ""
	} else if operationType == 3 {
		err := user.SetPassword(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByID(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	err = userDao.UpdateUserByID(userId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// Show 展示用户金额
func (service *Service) Show(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByID(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}
