package serializer

import (
	"TTMS_Web/conf"
	"TTMS_Web/model"
)

type Product struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	CreateAt      int64  `json:"create_at"`
	OnSale        bool   `json:"on_sale"`
	Num           int    `json:"num"`
	BossID        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(item *model.Product) Product {
	return Product{
		Id:            item.ID,
		Name:          item.Name,
		CategoryId:    item.CategoryId,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       conf.Config_.Path.Host + conf.Config_.Service.HttpPort + conf.Config_.Path.ProductPath + item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		CreateAt:      item.CreatedAt.Unix(),
		OnSale:        item.OnSale,
		Num:           item.Num,
		BossID:        item.BossID,
		BossName:      item.BossName,
		BossAvatar:    conf.Config_.Path.Host + conf.Config_.Service.HttpPort + conf.Config_.Path.AvatarPath + item.BossAvatar,
	}
}
