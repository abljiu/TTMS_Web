package dao

import (
	"TTMS_Web/model"
	"fmt"
)

// 自动迁移数据
func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.Actor{},
			&model.Theater{},
			&model.Admin{},
			&model.BasePage{},
			&model.Session{},
			&model.Hall{},
			&model.Director{},
			&model.Category{},
			&model.Notice{},
			&model.Order{},
			&model.Movie{},
			&model.User{},
			&model.Carousel{},
			&model.Comment{},
			&model.Upvote{},
		)
	if err != nil {
		fmt.Println("err:", err)
	}
	return
}
