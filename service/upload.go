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

// UploadMovieIndexToLocalStatic   更新电影封面图片到本地
func UploadMovieIndexToLocalStatic(file multipart.File, productName string) (filePath string, er error) {
	basePath := "." + conf.Config_.Path.MoviePath + productName + "/"
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
	return productName + "/" + productName + "index.jpg", err
}

// UploadDirectorToLocalStatic 更新导演图片到本地
func UploadDirectorToLocalStatic(files []*multipart.FileHeader, directors []string) (string, error) {
	var err error
	var directorPath string
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	basePath := "." + conf.Config_.Path.DirectorPath + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	for i := 0; i < len(files) && i < len(directors); i++ {
		director := directors[i]
		directorImg, _ := files[i].Open()
		directorPath = basePath + director + ".jpg"
		content, err := io.ReadAll(directorImg)
		if err != nil {
			return "", err
		}
		err = os.WriteFile(directorPath, content, 0666)
		if err != nil {
			return "", err
		}
	}
	return directorPath, err
}

// UploadActorToLocalStatic 更新演员图片到本地
func UploadActorToLocalStatic(files []*multipart.FileHeader, actors []string) (string, error) {
	var err error
	var actorPath string
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	basePath := "." + conf.Config_.Path.ActorPath + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	for i := 0; i < len(files) && i < len(actors); i++ {
		actor := actors[i]
		directorImg, _ := files[i].Open()
		actorPath = basePath + actor + ".jpg"
		content, err := io.ReadAll(directorImg)
		if err != nil {
			return "", err
		}
		err = os.WriteFile(actorPath, content, 0666)
		if err != nil {
			return "", err
		}
	}
	return actorPath, err
}

// UploadMovieToLocalStatic  更新电影图片到本地
func UploadMovieToLocalStatic(files []*multipart.FileHeader, productName string) (string, error) {
	var err error
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	basePath := "." + conf.Config_.Path.MoviePath + productName + "/"
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
	return productPath, err

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
