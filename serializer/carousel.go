package serializer

import (
	"TTMS_Web/conf"
	"TTMS_Web/model"
)

type Carousel struct {
	ImgPath   string `json:"img_path"`
	MovieName string `json:"movie_name"`
	Duration  string `json:"duration"`
}

func BuildCarousel(item *model.Carousel) Carousel {
	return Carousel{
		ImgPath:   conf.Config_.Path.Host + conf.Config_.Service.HttpPort + conf.Config_.Path.MoviePath + item.Movie.ImgPath,
		MovieName: item.Movie.ChineseName,
		Duration:  item.Movie.Duration.String(),
	}
}

func BuildCarousels(items []model.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousel := BuildCarousel(&item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
