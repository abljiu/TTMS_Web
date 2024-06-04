package serializer

import (
	"TTMS_Web/conf"
	"TTMS_Web/model"
	"strconv"
	"strings"
	"time"
)

type Movie struct {
	Id           uint             `json:"id"`
	ChineseName  string           `json:"chinese_name" `
	EnglishName  string           `json:"english_name" `
	CategoryID   []uint           `json:"category_id" `
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

func BuildMovie(item *model.Movie) Movie {
	strSlice := strings.Split(item.CategoryId, ",")
	numbers := make([]uint, len(strSlice))
	for i, s := range strSlice {
		num, err := strconv.ParseUint(s, 10, 0)
		if err != nil {
			continue
		}
		numbers[i] = uint(num)
	}
	return Movie{
		Id:           item.ID,
		ChineseName:  item.ChineseName,
		EnglishName:  item.EnglishName,
		CategoryID:   numbers,
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

func BuildMovies(items []*model.Movie) (products []Movie) {
	for i := 0; i < len(items); i++ {
		product := BuildMovie(items[i])
		products = append(products, product)
	}
	return products
}

func BuildMoviesByTheater(items []*model.MovieTheater) (products []Movie) {
	for i := 0; i < len(items); i++ {
		product := BuildMovie(&items[i].Movie)
		products = append(products, product)
	}
	return products
}
