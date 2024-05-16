package service

import (
	"TTMS_Web/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"sync"
)

// UploadAvatarToLocalStatic 更新头像到本地
func UploadAvatarToLocalStatic(file multipart.File, uid uint, userID string) (filePath string, er error) {
	bid := strconv.Itoa(int(uid))
	basePath := "." + conf.Config_.Path.AvatarPath + "user" + bid + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userID + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "user" + bid + "/" + userID + ".jpg", err
}

// UploadProductIndexToLocalStatic   更新电影封面图片到本地
func UploadProductIndexToLocalStatic(file multipart.File, productName string) (filePath string, er error) {
	basePath := "." + conf.Config_.Path.ProductPath + productName + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + "index.jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(productPath, content, 0666)
	if err != nil {
		return
	}
	return productPath + "/" + productName + ".jpg", err
}

// UploadProductToLocalStatic  更新电影图片到本地
func UploadProductToLocalStatic(files []*multipart.FileHeader, productName string) (string, error) {
	var err error
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	basePath := "." + conf.Config_.Path.ProductPath + productName + "/"
	var productPath string

	for num, file := range files {
		num := strconv.Itoa(num)
		tmp, _ := file.Open()
		productPath = basePath + productName + "_" + num + ".jpg"

		content, err := io.ReadAll(tmp)
		if err != nil {
			return "", err
		}
		err = os.WriteFile(productPath, content, 0666)
		if err != nil {
			return "", err
		}
	}
	return productPath + "/" + productName + ".jpg", err

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
