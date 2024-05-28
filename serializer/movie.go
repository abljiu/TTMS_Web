package serializer

import (
	"TTMS_Web/conf"
	"TTMS_Web/model"
	"time"
)

type Movie struct {
	Id           uint             `json:"id"`
	ChineseName  string           `json:"chinese_name" `
	EnglishName  string           `json:"english_name" `
	Category     string           `json:"category" `
	Area         string           `json:"area" `
	Duration     time.Duration    `json:"duration" `
	Showtime     time.Time        `json:"showtime"`
	Introduction string           `json:"introduction"`
	ImgPath      string           `json:"img_path"`
	OnSale       bool             `json:"on_sale"`
	Score        float64          `json:"score"`
	Sales        int64            `json:"sales"`
	Directors    []model.Director `json:"directors"`
	Actors       []model.Actor    `json:"actors"`
}

func BuildMovie(item *model.Movie, categoryString string) Movie {
	return Movie{
		Id:           item.ID,
		ChineseName:  item.ChineseName,
		EnglishName:  item.EnglishName,
		Category:     categoryString,
		Area:         item.Area,
		Duration:     item.Duration,
		Showtime:     item.ShowTime,
		Introduction: item.Introduction,
		ImgPath:      conf.Config_.Path.Host + conf.Config_.Service.HttpPort + conf.Config_.Path.MoviePath + item.ImgPath,
		OnSale:       item.OnSale,
		Score:        item.Score,
		Directors:    item.Directors,
		Actors:       item.Actors,
	}
}

func BuildMovies(items []*model.Movie, categoryStrings []string) (products []Movie) {
	for i := 0; i < len(items) && i < len(categoryStrings); i++ {
		product := BuildMovie(items[i], categoryStrings[i])
		products = append(products, product)
	}
	return products
}
