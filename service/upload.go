package service

import (
	"TTMS_Web/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

// UploadAvatarToLocalStatic 更新图像到本地
func UploadAvatarToLocalStatic(file multipart.File, uid uint, userName string) (filePath string, er error) {
	bid := strconv.Itoa(int(uid))
	basePath := "." + conf.Config_.Path.AvatarPath + "user" + bid + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "user" + bid + "/" + userName + ".jpg", err
}

// UploadProductToLocalStatic  更新图像到本地
func UploadProductToLocalStatic(file multipart.File, uid uint, productName string) (filePath string, er error) {
	bid := strconv.Itoa(int(uid))
	basePath := "." + conf.Config_.Path.ProductPath + "boss" + bid + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(productPath, content, 0666)
	if err != nil {
		return
	}
	return "boss" + bid + "/" + productName + ".jpg", err
}

// DirExistOrNot 判断路径是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}
