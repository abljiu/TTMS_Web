package service

import (
	"TTMS_Web/dao"
	"TTMS_Web/model"
	"TTMS_Web/pkg/e"
	"TTMS_Web/pkg/util"
	"TTMS_Web/serializer"
	"context"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

func (service *ProductService) Create(ctx context.Context, uid uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserByID(uid)
	//以第一张作为封面图
	tmp, _ := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, uid, service.Name)
	if err != nil {
		code = e.ErrorProductImgUpload
		util.LogrusObj.Infoln("UploadProductToLocalStatic", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	product := &model.Product{
		Name:          service.Name,
		CategoryId:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       service.ImgPath,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossID:        uid,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("CreateProduct", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = UploadProductToLocalStatic(tmp, uid, service.Name+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		productImg := model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
		asdasdasdasdasdasdasd
	}
}
