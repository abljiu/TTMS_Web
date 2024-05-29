package e

var MsgFlags = map[int]string{
	Success:                "ok",
	Error:                  "fail",
	InvalidParams:          "参数错误",
	ErrorNoPassword:        "密码不能为空",
	ErrorFailEncryption:    "密码加密失败",
	ErrorExistUserNotFound: "用户不存在",
	ErrorNotCompare:        "密码错误",
	ErrorAuthToken:         "token错误",
	ErrorAuthTokenTimeout:  "token过期",
	ErrorUploadFail:        "图片上传失败",
	ErrorSendEmail:         "发送邮件失败",
	ErrorStatus:            "您并非管理员",
	ErrorProductImgUpload:  "图片上传错误",
	ErrorMovieIndex:        "请上传电影封面",
	ErrorMovieId:           "电影id有误",
	ErrorHallId:            "影厅id有误",
	ErrorSessionId:         "场次id有误",
	ErrorInitializeStock:   "初始化库存失败",
	ErrorInvalidSeatParam:  "无效的座位参数",
}

// GetMsg 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
